package gw3

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
)

const (
	EmptyDAGRoot = "QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn"
)

func (c *Client) DAGAdd(root, path string, data []byte) (string, error) {
	url, err := c.AuthDAGAdd(root, path, len(data))
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	return resp.Header.Get("IPFS-Hash"), nil
}

func (c *Client) AuthDAGAdd(root, filePath string, size int) (string, error) {
	p := path.Join(root, filePath)
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf(
			"%s/ipfs/%s?size=%d",
			c.endPoint,
			p,
			size,
		),
		nil,
	)
	if err != nil {
		return "", err
	}

	var r redirect
	return r.URL, c.callGateway(req, &r)
}

func (c *Client) AuthDAGRemove(root, filePath string) (string, error) {
	p := path.Join(root, filePath)
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/ipfs/%s", c.endPoint, p),
		nil,
	)
	if err != nil {
		return "", err
	}

	var r redirect
	return r.URL, c.callGateway(req, &r)
}

func (c *Client) DAGRemove(root, path string) (string, error) {
	url, err := c.AuthDAGRemove(root, path)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	return resp.Header.Get("IPFS-Hash"), nil
}
