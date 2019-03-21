package main

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeHandle() *kubernetes.Clientset {
	var conf *rest.Config

	// In cluster config
	/*
		conf, err := rest.InClusterConfig()
		if err != nil {
			fmt.Println("err=", err)
		}*/

	// Outside of cluster config
	conf, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		fatal(fmt.Sprintf("error in getting Kubeconfig: %v", err))
	}

	cs, err := kubernetes.NewForConfig(conf)
	if err != nil {
		fatal(fmt.Sprintf("error in getting clientset from Kubeconfig: %v", err))
	}

	return cs
}

func fatal(msg string) {
	os.Stderr.WriteString(msg + "\n")
	os.Exit(1)
}
