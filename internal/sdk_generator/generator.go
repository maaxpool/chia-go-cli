package main

import (
	"chia-go-cli/sdk/client"
	"chia-go-cli/sdk/full_node"
	"chia-go-cli/sdk/shared"
	"chia-go-cli/sdk/wallet"
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	var err error
	err = ParseAndWrite(full_node.RpcApis, "full_node", "sdk/full_node/sdk_generator.go")
	if err != nil {
		panic(err)
	}
	err = ParseAndWrite(shared.RpcApis, "shared", "sdk/shared/sdk_generator.go")
	if err != nil {
		panic(err)
	}
	err = ParseAndWrite(wallet.RpcApis, "wallet", "sdk/wallet/sdk_generator.go")
	if err != nil {
		panic(err)
	}
}

func ParseAndWrite(method []*client.TemplateRpcMethod, namespace, path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0655)
	if err != nil {
		return err
	}

	_, _ = file.WriteString("// Code generated - DO NOT EDIT.\n// This file is a generated binding and any manual changes will be lost.\n")
	_, _ = file.WriteString(fmt.Sprintf("package %s\n", namespace))
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString("import (\n")
	_, _ = file.WriteString("\t\"chia-go-cli/sdk/client\"\n")
	_, _ = file.WriteString("\t\"encoding/json\"\n")
	_, _ = file.WriteString(")\n")
	_, _ = file.WriteString("\n")

	for i, rpcMethod := range method {
		methodString := genMethod(i, rpcMethod)
		_, err := file.WriteString(methodString)
		if err != nil {
			return err
		}
	}
	return nil
}

func genMethod(index int, rpcMethod *client.TemplateRpcMethod) string {
	code := make([]string, 0)
	funcParams := make([]string, 0)
	bodyParams := make([]string, 0)

	for i, val := range rpcMethod.ValInfo {
		lowerCamel := strcase.ToLowerCamel(val.Name)
		switch val.Type {
		case reflect.Int:
			funcParams = append(funcParams, fmt.Sprintf("%s int", lowerCamel))
			bodyParams = append(bodyParams, fmt.Sprintf("\treq.ValInfo[%d].Data = %s", i, lowerCamel))
		case reflect.String:
			funcParams = append(funcParams, fmt.Sprintf("%s string", lowerCamel))
			bodyParams = append(bodyParams, fmt.Sprintf("\treq.ValInfo[%d].Data = %s", i, lowerCamel))
		case reflect.Bool:
			funcParams = append(funcParams, fmt.Sprintf("%s bool", lowerCamel))
			bodyParams = append(bodyParams, fmt.Sprintf("\treq.ValInfo[%d].Data = %s", i, lowerCamel))
		default:
			log.Panicf("generator api command fail, not support type: %v", val.Type)
		}

		if val.FormatFunc != nil {
			bodyParams = append(bodyParams, fmt.Sprintf("\treq.ValInfo[%d].FormatFunc = RpcApis[%d].ValInfo[%d].FormatFunc", i, index, i))
		}
	}

	requestRaw, err := json.Marshal(rpcMethod)
	if err != nil {
		panic(err)
	}

	code = append(code, fmt.Sprintf("// %s %s", rpcMethod.MethodName, rpcMethod.Desc))
	code = append(code, fmt.Sprintf("func %s(%s) client.RpcMethod {", rpcMethod.MethodName, strings.Join(funcParams, ", ")))
	code = append(code, fmt.Sprintf("\treq := &client.TemplateRpcMethod{}"))
	code = append(code, fmt.Sprintf("\t_ = json.Unmarshal([]byte(%s), req)", strconv.Quote(string(requestRaw))))
	code = append(code, bodyParams...)
	code = append(code, fmt.Sprintf("\treturn req"))
	code = append(code, fmt.Sprintf("}"))
	code = append(code, "")
	code = append(code, "")

	return strings.Join(code, "\n")
}
