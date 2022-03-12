/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dummy

import (
	"context"
	"fmt"
	"sync"
	"time"

	"k8s.io/klog/v2"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	clientretry "k8s.io/client-go/util/retry"
	"k8s.io/component-base/metrics/prometheus/ratelimiter"
	nodeutil "k8s.io/component-helpers/node/util"
)

var updateNetworkConditionBackoff = wait.Backoff{
	Steps:    5, // Maximum number of retries.
	Duration: 100 * time.Millisecond,
	Jitter:   1.0,
}

type RouteController struct {
	kubeClient       clientset.Interface
	clusterName      string
	nodeLister       corelisters.NodeLister
	nodeListerSynced cache.InformerSynced
	broadcaster      record.EventBroadcaster
	recorder         record.EventRecorder
}

func New(kubeClient clientset.Interface, nodeInformer coreinformers.NodeInformer, clusterName string) *RouteController {
	if kubeClient != nil && kubeClient.CoreV1().RESTClient().GetRateLimiter() != nil {
		ratelimiter.RegisterMetricAndTrackRateLimiterUsage("dummy_controller", kubeClient.CoreV1().RESTClient().GetRateLimiter())
	}

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "dummy_controller"})

	rc := &RouteController{
		kubeClient:       kubeClient,
		clusterName:      clusterName,
		nodeLister:       nodeInformer.Lister(),
		nodeListerSynced: nodeInformer.Informer().HasSynced,
		broadcaster:      eventBroadcaster,
		recorder:         recorder,
	}

	return rc
}

func (rc *RouteController) Run(ctx context.Context, syncPeriod time.Duration) {
	defer utilruntime.HandleCrash()

	klog.Info("Starting dummy controller")
	defer klog.Info("Shutting down dummy controller")

	if rc.broadcaster != nil {
		rc.broadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: rc.kubeClient.CoreV1().Events("")})
	}

	// TODO: If we do just the full Resync every 5 minutes (default value)
	// that means that we may wait up to 5 minutes before even starting
	// creating a route for it. This is bad.
	// We should have a watch on node and if we observe a new node (with CIDR?)
	// trigger reconciliation for that node.
	go wait.NonSlidingUntil(func() {
		if err := rc.reconcileNodes(ctx); err != nil {
			klog.Errorf("Couldn't reconcile nodes: %v", err)
		}
	}, syncPeriod, ctx.Done())

	<-ctx.Done()
}

func (rc *RouteController) reconcileNodes(ctx context.Context) error {
	nodes, err := rc.nodeLister.List(labels.Everything())
	if err != nil {
		return fmt.Errorf("error listing nodes: %v", err)
	}
	return rc.reconcile(ctx, nodes)
}

func (rc *RouteController) reconcile(ctx context.Context, nodes []*v1.Node) error {
	wg := sync.WaitGroup{}

	for _, node := range nodes {
		wg.Add(1)
		go func(n *v1.Node) {
			defer wg.Done()
			rc.updateNetworkingCondition(n)
		}(node)
	}

	wg.Wait()
	return nil
}

func (rc *RouteController) updateNetworkingCondition(node *v1.Node) error {
	_, condition := nodeutil.GetNodeCondition(&(node.Status), v1.NodeNetworkUnavailable)
	if condition != nil && condition.Status == v1.ConditionFalse {
		return nil
	}

	err := clientretry.RetryOnConflict(updateNetworkConditionBackoff, func() error {
		var err error
		// Patch could also fail, even though the chance is very slim. So we still do
		// patch in the retry loop.
		currentTime := metav1.Now()
		err = nodeutil.SetNodeCondition(rc.kubeClient, types.NodeName(node.Name), v1.NodeCondition{
			Type:               v1.NodeNetworkUnavailable,
			Status:             v1.ConditionFalse,
			Reason:             "NodeConfigured",
			Message:            "DummyController removed default taints",
			LastTransitionTime: currentTime,
		})
		if err != nil {
			klog.V(4).Infof("Error updating node %s, retrying: %v", types.NodeName(node.Name), err)
		}
		return err
	})
	if err != nil {
		klog.Errorf("Error updating node %s: %v", node.Name, err)
	}

	return err
}
