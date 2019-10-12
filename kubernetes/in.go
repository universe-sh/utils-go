package kubernetes

import (
	"github.com/hashicorp/go-hclog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	// Auth K8s on all Cloud Providers
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// NewClient returns a k8s clientset to the request from inside of cluster
func NewClient(logger hclog.Logger) kubernetes.Interface {
	clientset, err := kubernetes.NewForConfig(BuildInClusterConfig(logger))
	if err != nil {
		return nil
	}

	return clientset
}

// BuildInClusterConfig Configuration in AC Kubernetes
func BuildInClusterConfig(logger hclog.Logger) *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil
	}

	return config
}
