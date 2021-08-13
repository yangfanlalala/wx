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
	Method string `json:"-"`
	ContentType string `json:"-"`
	URL string `json:"-"`
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

func (cr *CommonRequest) BuildRequest() (*http.Request, error)  {
	fmt.Println(cr)
	val := reflect.ValueOf(cr)
	typ := val.Type()
	u, err := url.Parse(cr.URL)
	if err != nil {
		return nil, err
	}
	for i :=0; i < typ.Elem().NumField(); i++ {
		//typ.Field(i)
		//if typ.Field(i).Tag.Get("position") == "query" {
		//	//fn := typ.Field(i).Tag.Get("name")
		//	//if fn == "" {
		//	//	fn = typ.Field(i).Name
		//	//}
		//	//u.Query().Add(fn, val.Field(i).String())
		//}
	}
	buf := bytes.Buffer{}
	if err = json.NewEncoder(&buf).Encode(cr); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(cr.Method, u.String(), &buf)
	return req, err
}
