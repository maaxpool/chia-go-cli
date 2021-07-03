package common

import (
	"bytes"
	"github.com/tidwall/sjson"
	"net/http"
	"net/url"
	"path"
	"reflect"
)

type TemplateValue struct {
	Name    string
	Desc    string
	Type    reflect.Kind
	Default interface{}
	Path    string
	Data    interface{}
}

type TemplateRpcMethod struct {
	MethodName   string
	Desc         string
	Method       string
	Url          string
	JsonTemplate string
	ValInfo      []*TemplateValue
}

func (t *TemplateRpcMethod) Name() string {
	return t.MethodName
}

func (t *TemplateRpcMethod) BuildRequest(url *url.URL) (*http.Request, error) {
	body, err := t.buildBody()
	if err != nil {
		return nil, err
	}

	url.Path = path.Join(url.Path, t.Url)
	return http.NewRequest(t.Method, url.String(), bytes.NewReader(body))
}

func (t *TemplateRpcMethod) buildBody() (ret []byte, err error) {
	if t.JsonTemplate == "" {
		return []byte{}, nil
	}

	template := t.JsonTemplate
	for _, val := range t.ValInfo {
		template, err = sjson.Set(template, val.Path, val.Data)
		if err != nil {
			return nil, err
		}
	}

	return []byte(template), nil
}
