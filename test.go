package main

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"fmt"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "", "Path to a kubeconfig file")

	flag.Parse()

	if *kubeconfig == "" {
		panic("no kubeconfig provided")
	} else {
		fmt.Println("Using kubeconfig:", *kubeconfig)
	}

	clientSet := newClientSet(*kubeconfig)

	testing(clientSet)
}

func testing(clientset kubernetes.Interface) {
	//TODO
}

func newClientSet(kubeconfig string) kubernetes.Interface {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	return clientset
}
