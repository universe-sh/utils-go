package httputils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"
)

// URLQuery HTTP
func URLQuery(queries url.Values) map[string][]string {
	var parameters = make(map[string][]string)

	for name, values := range queries {
		parameters[name] = values
	}

	return parameters
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
