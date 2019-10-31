package kubernetes

import (
	"fmt"
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
	var (
		clientset *kubernetes.Clientset = nil
		err       error                 = nil
	)

	if clientset, err = kubernetes.NewForConfig(BuildOutOfClusterConfig(logger)); err != nil {
		logger.With("error", err, "context", "out").Error("kubernetes config")
		os.Exit(1)
	}

	return clientset
}

// BuildOutOfClusterConfig func
func BuildOutOfClusterConfig(logger hclog.Logger) *rest.Config {
	var (
		kubeconfigPath string       = os.Getenv("KUBECONFIG")
		home           string       = os.Getenv("HOME")
		config         *rest.Config = nil
		err            error        = nil
	)

	if kubeconfigPath == "" {
		kubeconfigPath = fmt.Sprintf("%s/.kube/config", home)
	}

	if config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath); err != nil {
		logger.With("error", err, "context", "out").Error("kubernetes config")
		os.Exit(1)
	}

	return config
}
