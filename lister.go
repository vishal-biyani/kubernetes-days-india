package main

import (
	"fmt"

	"github.com/urfave/cli"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
)

func lister(c *cli.Context) {
	fmt.Println("Running Lister Example")
	cs := getKubeHandle()
	listWatch := cache.NewListWatchFromClient(cs.Core().RESTClient(), "pods", "", fields.Everything())

	ro, err := listWatch.List(metav1.ListOptions{})
	if err != nil {
		fmt.Println("err=", err)
	}

	pods := ro.(*v1.PodList)

	fmt.Println("Pods from Lister")
	for j, pod := range pods.Items {
		fmt.Printf("%d) %v \n", j, pod.Name)
	}

}
