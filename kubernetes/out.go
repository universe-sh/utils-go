package kubernetes

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	// Auth K8s on all Cloud Providers
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// NewClientOutOfCluster returns a k8s clientset to the request from outside of cluster
func NewClientOutOfCluster(logger hclog.Logger) kubernetes.Interface {
	clientset, err := kubernetes.NewForConfig(BuildOutOfClusterConfig(logger))
	if err != nil {
		logger.With("error", err).Error("Kubernetes config")
		os.Exit(1)
	}

	return clientset
}

// BuildOutOfClusterConfig func
func BuildOutOfClusterConfig(logger hclog.Logger) *rest.Config {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		kubeconfigPath = os.Getenv("HOME") + "/.kube/config"
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		logger.With("error", err).Error("Kubernetes config")
		os.Exit(1)
	}

	return config
}
