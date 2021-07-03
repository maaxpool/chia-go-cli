package common

import (
	"bytes"
	"net/http"
	"net/url"
	"path"
)

type RpcMethod interface {
	Name() string
	BuildRequest(url *url.URL) (*http.Request, error)
}

type RawRpcMethod struct {
	MethodName string
	Method     string
	Url        string
	Body       []byte
}

func NewRawRpcMethod(method string, url string, body []byte) *RawRpcMethod {
	return &RawRpcMethod{MethodName: "raw", Method: method, Url: url, Body: body}
}

func (r *RawRpcMethod) Name() string {
	return r.MethodName
}

func (r *RawRpcMethod) BuildRequest(url *url.URL) (*http.Request, error) {
	url.Path = path.Join(url.Path, r.Url)
	return http.NewRequest(r.Method, url.String(), bytes.NewReader(r.Body))
}
