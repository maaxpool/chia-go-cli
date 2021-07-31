package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/sjson"
	"net/http"
	"net/url"
	"path"
	"reflect"
)

type TemplateValue struct {
	Name       string
	Desc       string
	Type       reflect.Kind
	Default    interface{}
	Path       string
	Data       interface{}
	FormatFunc func(data interface{}) (interface{}, error) `json:"-"`
}

type TemplateRpcMethod struct {
	MethodName   string
	Desc         string
	Method       string
	Url          string
	JsonTemplate string
	ValInfo      []*TemplateValue
}

var String2JsonFormatFunc = func(data interface{}) (interface{}, error) {
	var (
		result  map[string]interface{}
		jsonRaw string
		ok      bool
	)

	if jsonRaw, ok = data.(string); !ok {
		return nil, errors.New("can not convert param to string")
	}

	err := json.Unmarshal([]byte(jsonRaw), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
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
		refData := reflect.ValueOf(val.Data)
		if refData.IsNil() {
			return nil, fmt.Errorf("value %s is nil", val.Name)
		}

		if refData.Kind() == reflect.Ptr {
			refData = refData.Elem()
		}

		data := refData.Interface()
		if val.FormatFunc != nil {
			data, err = val.FormatFunc(refData.Interface())
			if err != nil {
				return nil, err
			}
		}

		template, err = sjson.Set(template, val.Path, data)
		if err != nil {
			return nil, err
		}
	}

	return []byte(template), nil
}
