package httputils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/universe-sh/utils-go/slice"
)

// Pages Pagination Limits
type Pages struct {
	Start int
	End   int
}

// Pagination struct
type Pagination struct {
	PerPage int
	Page    int
	Limits  Pages
}

// Data result reponse
type Data struct {
	Count   int           `json:"count"`
	Results []interface{} `json:"results"`
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
		paginate   Pagination
		value      int
		err        error
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
		}

		parameters[name] = values
	}

	// Pagination calc limitations
	if paginate.PerPage >= minPerPage && paginate.PerPage <= maxPerPage {
		paginate.Limits.Start = ((paginate.Page * paginate.PerPage) - paginate.PerPage)
		paginate.Limits.End = (paginate.Page * paginate.PerPage) - 1
	} else {
		paginate = Pagination{
			PerPage: 25, Page: 1,
			Limits: Pages{Start: 0, End: 24},
		}
	}

	return parameters, paginate
}

// PostJSONQuery HTTP
func PostJSONQuery(rbody io.ReadCloser) interface{} {
	var data interface{}

	// POST Body
	body, err := ioutil.ReadAll(rbody)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}

	return data
}

// DataQuery informations
func DataQuery(tmpData *Data, pages Pagination) *Data {
	var (
		count = tmpData.Count - 1
		data  = new(Data)
	)

	for i := pages.Limits.Start; i <= pages.Limits.End; i++ {
		if i > count {
			break
		}

		data.Results = append(data.Results, tmpData.Results[i])
		data.Count++
	}

	return data
}
