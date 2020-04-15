package main

import (
	"flag"
	"log"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/dynamic"
)

func GetDynamicClientOrFail() (dynamic.Interface) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	dynClient, errClient := dynamic.NewForConfig(config)
	if errClient != nil {
		log.Fatalf("Error received creating client %v", errClient)
	}
	return dynClient
}