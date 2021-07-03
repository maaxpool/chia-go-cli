package common

import (
	"chia-go-cli/logic/node"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
	typ        node.CertType
	node       *node.Node
	port       uint16
}

func NewClient(node *node.Node, port uint16, typ node.CertType) (*Client, error) {
	cert, err := node.GetCert(typ)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{*cert},
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &Client{
		node:       node,
		port:       port,
		typ:        typ,
		httpClient: &http.Client{Transport: transport},
	}, nil
}

func (c *Client) RawRequest(method RpcMethod) (*http.Response, error) {
	hostUrl, err := url.Parse(c.node.NodeUrl())
	if err != nil {
		return nil, err
	}

	hostUrl.Host = fmt.Sprintf("%s:%d", hostUrl.Hostname(), c.port)

	req, err := method.BuildRequest(hostUrl)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	return c.httpClient.Do(req)
}
