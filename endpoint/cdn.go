package endpoint

import (
	"fmt"
)

// Region struct
type Region struct {
	Name  string   `json:"name,omitempty"`
	Zones []string `json:"zones,omitempty"`
}

// Quotas Resources
type Quotas struct {
	CPU    int32 `json:"cpu,omitempty"`
	Memory int32 `json:"ram,omitempty"`
	GPU    int32 `json:"gpu,omitempty"`
}

// Kubernetes Quotas
type Kubernetes struct {
	Requests *Quotas `json:"requests,omitempty"`
	Limits   *Quotas `json:"limits,omitempty"`
}

// Resources struct
type Resources struct {
	Compute    *Quotas     `json:"compute,omitempty"`
	Kubernetes *Kubernetes `json:"kubernetes,omitempty"`
}

// Pricing struct
type Pricing struct {
	OnDemand float32            `json:"ondemand,omitempty"`
	Spot     map[string]float32 `json:"spot,omitempty"`
}

// Engine struct
type Engine struct {
	Resources *Resources `json:"resources,omitempty"`
	Pricing   *Pricing   `json:"pricing,omitempty"`
}

// GetRegions Cloud Provider
func (c *Client) GetRegions(cloud string) ([]Region, error) {
	var (
		r   []Region
		err error
	)

	url := fmt.Sprintf("https://%s/dyn/%s/regions.json", cdn, cloud)
	err = c.GetJSON(url, &r)

	return r, err
}

// GetEngine Cloud Provider by region
func (c *Client) GetEngine(cloud, region string) (map[string]Engine, error) {
	var (
		r   = make(map[string]map[string]Engine, 0)
		err error
	)

	url := fmt.Sprintf("https://%s/dyn/%s/engines.json", cdn, cloud)
	err = c.GetJSON(url, &r)

	return r[region], err
}
