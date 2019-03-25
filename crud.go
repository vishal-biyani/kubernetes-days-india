package main

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/urfave/cli"
)

func crudOperation(c *cli.Context) {
	fmt.Println("Running CRUD Example")
	cs := getKubeHandle()

	pods, err := cs.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		fatal(fmt.Sprintf("error getting list of pods: %v", err))
	}

	fmt.Println("## Pods ##")
	for i, pod := range pods.Items {
		fmt.Printf("%d) %v \n", i, pod.Name)
	}
}
