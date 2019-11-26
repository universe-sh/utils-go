package httputils

import (
	"io"
	"io/ioutil"
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
		} else {
			parameters[name] = values
		}
	}

	paginate.Limits = make(map[string]int)
	// Pagination calc limitations
	if paginate.PerPage >= minPerPage && paginate.PerPage <= maxPerPage {
		paginate.Limits["first"] = ((paginate.Page * paginate.PerPage) - paginate.PerPage)
		paginate.Limits["last"] = (paginate.Page * paginate.PerPage) - 1
	} else {
		paginate = Pagination{
			PerPage: 25, Page: 1,
			Limits: map[string]int{"first": 0, "last": 24},
		}
	}

	return parameters, paginate
}

func BodyQuery(body io.ReadCloser) ([]byte, error) {
	var (
		data []byte
		err  error
	)

	if data, err = ioutil.ReadAll(io.LimitReader(body, 1048576)); err != nil {
		return nil, err
	}

	if err = body.Close(); err != nil {
		return nil, err
	}

	return data, nil
}

// DataQuery informations
func DataQuery(tmpData []interface{}, pages Pagination) *Response {
	var (
		count = len(tmpData) - 1
		data  = make([]interface{}, 0)
		i     int
	)

	for i = pages.Limits["first"]; i <= pages.Limits["last"]; i++ {
		if i > count {
			break
		}

		data = append(data, tmpData[i])
	}

	return &Response{
		Results: data,
		Metadatas: &Metadatas{
			TotalIndex:     (count + 1),
			FirstIndexPage: (pages.Limits["first"] + 1),
			LastIndexPage:  i,
		},
	}
}
