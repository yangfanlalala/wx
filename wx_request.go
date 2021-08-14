package wx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

type Request interface {
	BuildRequest() (*http.Request, error)
}

type CommonRequest struct {
	Method      string
	ContentType string
	URL         string
	Data        interface{}
}

func (cr *CommonRequest) WithURL(url string) *CommonRequest {
	cr.URL = url
	return cr
}

func (cr *CommonRequest) WithMethod(method string) *CommonRequest {
	cr.Method = method
	return cr
}

func (cr *CommonRequest) WithContentType(ct string) *CommonRequest {
	cr.ContentType = ct
	return cr
}

func (cr *CommonRequest) WithData(data interface{}) *CommonRequest {
	cr.Data = data
	return cr
}

func (cr *CommonRequest) BuildRequest() (*http.Request, error) {
	val := reflect.ValueOf(cr.Data)
	typ := val.Type().Elem()
	u, err := url.Parse(cr.URL)
	if err != nil {
		return nil, err
	}
	values := u.Query()
	for i := 0; i < typ.NumField(); i++ {
		if t, _ := typ.Field(i).Tag.Lookup("position"); t == "query" {
			fn := typ.Field(i).Tag.Get("name")
			if fn == "" {
				fn = typ.Field(i).Name
			}
			fmt.Println(fn)
			if typ.Field(i).Type.Kind() == reflect.String {
				values.Add(fn, val.Elem().Field(i).String())
			}
		}
	}
	buf := bytes.Buffer{}
	if err = json.NewEncoder(&buf).Encode(cr); err != nil {
		return nil, err
	}
	u.RawQuery = values.Encode()
	fmt.Println(u.String())
	req, err := http.NewRequest(cr.Method, u.String(), &buf)
	return req, err
}
