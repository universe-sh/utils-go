package kubernetes

import (
	"github.com/hashicorp/go-hclog"
	"k8s.io/client-go/rest"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

// NewMetricsClient returns a k8s metrics client
func NewMetricsClient(c *rest.Config, logger hclog.Logger) *metrics.Clientset {
	mc, err := metrics.NewForConfig(c)
	if err != nil {
		logger.With("error", err).Error("Impossible to fetch metrics, server-metrics is installed ?")
		return nil
	}

	return mc
}
