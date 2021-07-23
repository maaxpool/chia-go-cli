package common

import (
	"chia-go-cli/logic/node"
	sdk "chia-go-cli/sdk/client"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	client *sdk.Client
}

func NewClient(node *node.Node, port uint16, typ node.CertType) (*Client, error) {
	cert, err := node.GetCert(typ)
	if err != nil {
		return nil, err
	}

	hostUrl, err := url.Parse(node.NodeUrl())
	if err != nil {
		return nil, err
	}

	hostUrl.Host = fmt.Sprintf("%s:%d", hostUrl.Hostname(), port)

	client, err := sdk.NewClient(hostUrl.String(), cert)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) RawRequest(method sdk.RpcMethod) (*http.Response, error) {
	return c.client.RawRequest(method)
}
