package endpoint

import (
	"fmt"
)

// Quotas Resources
type Quotas struct {
	CPU     int64 `json:"cpu,omitempty"`
	GPU     int64 `json:"gpu,omitempty"`
	Memory  int64 `json:"memory,omitempty"`
	Pods    int64 `json:"pods,omitempty"`
	Network int64 `json:"network,omitempty"`
}

// Kubernetes Quotas
type Kubernetes struct {
	Request Quotas `json:"requests,omitempty"`
	Limits  Quotas `json:"limits,omitempty"`
}

// Resources cpu & memory
type Resources struct {
	Compute    Quotas     `json:"compute,omitempty"`
	Kubernetes Kubernetes `json:"kubernetes,omitempty"`
}

// Properties instance
type Properties struct {
	Resources Resources         `json:"resources,omitempty"`
	Regions   map[string]Prices `json:"regions,omitempty"`
}

// Prices instances
type Prices struct {
	OnDemand float64            `json:"ondemand,omitempty"`
	Spot     map[string]float64 `json:"spot,omitempty"`
}

// GetRegions Cloud Provider
func (c *Client) GetRegions(cloud string) (map[string][]string, error) {
	var (
		r   map[string][]string
		err error
	)

	url := fmt.Sprintf("https://%s/dyn/%s/regions.json", cdn, cloud)
	err = c.GetJSON(url, &r)

	return r, err
}

// GetEngine Cloud Provider by region
func (c *Client) GetEngine(cloud, region string) (map[string]*Properties, error) {
	var (
		r   = make(map[string]*Properties, 0)
		err error
	)

	url := fmt.Sprintf("https://%s/dyn/%s/engines.json", cdn, cloud)
	err = c.GetJSON(url, &r)

	return r, err
}
