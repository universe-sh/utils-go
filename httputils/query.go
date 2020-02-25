package httputils

import (
	"net/url"
	"strconv"

	"github.com/universe-sh/utils-go/slice"
)

// Pagination struct
type Pagination struct {
	PerPage int
	Page    int
	Limits  map[string]int
}

const (
	minPerPage = 1
	maxPerPage = 50
)

var (
	pageFlags = []string{"page", "per_page"}
)

// URLQuery HTTP
func URLQuery(queries url.Values) (map[string][]string, Pagination) {
	var (
		parameters = make(map[string][]string)
		paginate   = Pagination{
			PerPage: 25, Page: 1,
			Limits: map[string]int{
				"first": 0,
				"last":  24,
			},
		}
		value int
		err   error
	)

	for name, values := range queries {
		value = 0

		if slice.StringInSlice(name, pageFlags) {
			if len(values) == 1 {
				if value, err = strconv.Atoi(values[0]); err == nil {
					switch name {
					case "page":
						paginate.Page = value
					case "per_page":
						paginate.PerPage = value
					}
				}
			}
		} else {
			parameters[name] = values
		}
	}

	paginate.Limits = make(map[string]int)
	// Pagination calc limitations
	if paginate.PerPage >= minPerPage && paginate.PerPage <= maxPerPage {
		paginate.Limits["first"] = ((paginate.Page * paginate.PerPage) - paginate.PerPage)
		paginate.Limits["last"] = (paginate.Page * paginate.PerPage) - 1
	}

	return parameters, paginate
}

// DataQuery informations
func DataQuery(tmpData []interface{}, pages Pagination) *Response {
	var (
		count = len(tmpData)
		data  = make([]interface{}, 0)
		last  int
	)

	for last = pages.Limits["first"]; last <= pages.Limits["last"]; last++ {
		if last >= count {
			break
		}

		data = append(data, tmpData[last])
	}

	return &Response{
		Results: data,
		Metadatas: &Metadatas{
			TotalIndex:     count,
			FirstIndexPage: (pages.Limits["first"] + 1),
			LastIndexPage:  last,
		},
	}
}
