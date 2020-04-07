package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func configureClientset() *kubernetes.Clientset {
	homeDir := homedir.HomeDir()
	kubeconfig := filepath.Join(homeDir, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err)
	}

	return clientset
}

func combineConfigMaps(configmapList []*v1.ConfigMap) {
	fmt.Println(configmapList)
}

func configureConfigmapInformer(clientset *kubernetes.Clientset) cache.SharedIndexInformer {
	factory := informers.NewSharedInformerFactoryWithOptions(
		clientset,
		time.Second*60,
		informers.WithNamespace("default"),
		informers.WithTweakListOptions(func(listOptions *metav1.ListOptions) {
			listOptions.LabelSelector = "findme=foo"
		}),
	)

	configMapInformer := factory.Core().V1().ConfigMaps().Informer()
	configMapInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			configmap := obj.(*v1.ConfigMap)
			log.Printf("Configmap Added to store: %s", configmap.GetName())
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			configmap := oldObj.(*v1.ConfigMap)

			store := configMapInformer.GetStore()

			//dunno what to do here =/
			combineConfigMaps(store.List())

			log.Printf("Configmap updated in store: %s", configmap.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			configmap := obj.(*v1.ConfigMap)
			log.Printf("Configmap removed from store: %s", configmap.GetName())
		},
	})

	return configMapInformer
}

func main() {
	clientset := configureClientset()

	informer := configureConfigmapInformer(clientset)

	stopper := make(chan struct{})
	defer runtime.HandleCrash()
	defer close(stopper)

	informer.Run(stopper)
}
