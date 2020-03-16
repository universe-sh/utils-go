package types

var (
	provider = map[string]map[string]int64{
		"google": map[string]int64{
			"asia-east1":              1,
			"asia-east2":              2,
			"asia-northeast1":         3,
			"asia-northeast2":         4,
			"asia-south1":             5,
			"asia-southeast1":         6,
			"australia-southeast1":    7,
			"europe-north1":           8,
			"europe-west1":            9,
			"europe-west2":            10,
			"europe-west3":            11,
			"europe-west4":            12,
			"europe-west6":            13,
			"northamerica-northeast1": 14,
			"southamerica-east1":      15,
			"us-central1":             16,
			"us-east1":                17,
			"us-east4":                18,
			"us-west1":                19,
			"us-west2":                20,
		},
		"amazon": map[string]int64{},
	}
)

// StringtoInt64 converter
func StringtoInt64(p string) int64 {
	var i int64 = 1

	for k := range provider {
		if k == p {
			return int64(i)
		}

		i++
	}

	return int64(0)
}

// RegiontoString string converter
func RegiontoString(cloud, region int64) string {
	var i int64 = 1

	for _, v := range provider {
		if i == int64(cloud) {
			for k, j := range v {
				if j == region {
					return k
				}
			}
		}

		i++
	}

	return ""
}
