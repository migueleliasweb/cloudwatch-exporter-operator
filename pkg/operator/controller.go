package operator

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}

func combineConfigMaps(configmapList []*v1.ConfigMap) {
	fmt.Println(configmapList)
}

// Controller It controls...things...
type Controller struct {
	informer cache.SharedIndexInformer
	stopper  chan struct{}
}

func (controller *Controller) setup() {
	controller.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			configmap := obj.(*v1.ConfigMap)
			log.Debugf("Configmap Added to store: %s", configmap.GetName())
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			configmap := oldObj.(*v1.ConfigMap)

			// store := configMapInformer.GetStore()

			//dunno what to do here =/
			// combineConfigMaps(store.List())

			// configMapInformer.GetController()

			log.Debugf("Configmap updated in store: %s", configmap.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			configmap := obj.(*v1.ConfigMap)
			log.Debugf("Configmap removed from store: %s", configmap.GetName())
		},
	})
}

// Run Starts the controller
func (controller *Controller) Run() {
	defer runtime.HandleCrash()
	controller.informer.Run(controller.stopper)
}

// Stop Stops the controller by closing the run channel
func (controller *Controller) Stop() {
	close(controller.stopper)
}

// NewController Constructs a new Controller instance
func NewController(
	clientset kubernetes.Interface,
) *Controller {
	c := &Controller{
		informer: buildConfigMapInformer(
			clientset,
			"default",
			"findme=foo",
		),
		stopper: make(chan struct{}),
	}

	c.setup()

	return c
}
