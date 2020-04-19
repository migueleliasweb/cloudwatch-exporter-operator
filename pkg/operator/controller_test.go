package operator

import (
	"testing"
	"time"

	"k8s.io/client-go/kubernetes/fake"
)

// Interesting link: https://github.com/kubernetes/client-go/blob/master/examples/fake-client/main_test.go

func TestNewController(t *testing.T) {

	// Create the fake client.
	fakeClient := fake.NewSimpleClientset()

	controller := NewController(fakeClient)

	go func(c *Controller) {
		select {
		case <-time.After(time.Second * 1):
			controller.Stop()
		}
	}(controller)

	controller.Run()
}
