package main

import (
	"fmt"
	"time"

	"github.com/urfave/cli"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
)

func informer(c *cli.Context) {
	fmt.Println("Running Example Three")
	cs := getKubeHandle()
	listWatch := cache.NewListWatchFromClient(cs.Core().RESTClient(), "pods", "", fields.Everything())

	_, controller := cache.NewInformer(listWatch, &v1.Pod{}, time.Second*5, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			fmt.Println("Pod Added:", pod.Name)
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			fmt.Println("Pod Deleted:", pod.Name)
		},
	})

	stop := make(chan struct{})
	controller.Run(stop)
}
