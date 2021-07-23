package client

import (
	"chia-go-cli/logic/node"
	"crypto/tls"
	"net/http"
	"net/url"
)

type Client struct {
	host       string
	httpClient *http.Client
	typ        node.CertType
}

func NewClient(host string, cert *tls.Certificate) (*Client, error) {
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{*cert},
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &Client{
		host:       host,
		httpClient: &http.Client{Transport: transport},
	}, nil
}

func (c *Client) RawRequest(method RpcMethod) (*http.Response, error) {
	hostUrl, err := url.Parse(c.host)
	if err != nil {
		return nil, err
	}

	req, err := method.BuildRequest(hostUrl)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	return c.httpClient.Do(req)
}
