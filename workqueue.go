package main

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/urfave/cli"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
	informer cache.Controller
}

func workqueue_example(c *cli.Context) {
	fmt.Println("Running Four")
	cs := getKubeHandle()

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	listWatch := cache.NewListWatchFromClient(cs.Core().RESTClient(), "pods", "", fields.Everything())

	indexer, informer := cache.NewIndexerInformer(listWatch, &v1.Pod{}, time.Second*5, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	controller := &Controller{
		informer: informer,
		indexer:  indexer,
		queue:    queue,
	}

	stop := make(chan struct{})
	go controller.Run(stop)

	// Wait forever
	select {}
}

func (c *Controller) Run(stopCh chan struct{}) {
	defer runtime.HandleCrash()

	// Let the workers stop when we are done
	defer c.queue.ShutDown()

	go c.informer.Run(stopCh)

	// Wait for all involved caches to be synced, before processing items from the queue is started
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	go wait.Until(c.runWorker, time.Second, stopCh)
	<-stopCh
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}

	defer c.queue.Done(key)

	err := c.processBusinessLogic(key.(string))

	c.handleErr(err, key)

	return true
}

func (c *Controller) processBusinessLogic(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)

	if err != nil {
		glog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		// Below we will warm up our cache with a Pod, so that we will see a delete for one pod
		fmt.Printf("Pod %s does not exist anymore\n", key)
	} else {
		// Note that you also have to check the uid if you have a local controlled resource, which
		// is dependent on the actual instance, to detect that a Pod was recreated with the same name
		fmt.Printf("Sync/Add/Update for Pod %s\n", obj.(*v1.Pod).GetName())
	}

	return nil
}

func (c *Controller) handleErr(err error, key interface{}) {
	glog.Infof("Dropping pod %q out of the queue: %v", key, err)
}
