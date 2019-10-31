package kubernetes

import (
	"github.com/hashicorp/go-hclog"
	"k8s.io/client-go/rest"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

// NewMetricsClient returns a k8s metrics client
func NewMetricsClient(c *rest.Config, logger hclog.Logger) *metrics.Clientset {
	var (
		mc  *metrics.Clientset = nil
		err error              = nil
	)

	if mc, err = metrics.NewForConfig(c); err != nil {
		logger.With("error", err).Error("Impossible to fetch metrics, server-metrics is installed ?")
		return nil
	}

	return mc
}
