package main

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/golang/glog"
	"fmt"
)

const Port = ":8081"

func main() {
	kubeconfig := flag.String("kubeconfig", "", "Path to a kubeconfig file")

	flag.Parse()

	clientSet := newClientSet(*kubeconfig)

	podList, err := clientSet.CoreV1().Pods("default").List(metaV1.ListOptions{})

	if err != nil {
		panic(err)
	}

	for _,pod := range podList.Items {
		fmt.Println(pod.Name)
	}
}

// returns config using kubeconfig if provided, else from cluster context
func newClientSet(kubeconfig string) kubernetes.Interface {

	var config *rest.Config
	var err error

	if kubeconfig != "" {
		glog.Infof("using provided kubeconfig")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		glog.Infof("no kubeconfig provided, assuming we are running inside a cluster")
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	return clientset
}
