package gw3

import (
	"bytes"
	"fmt"
	"net/http"
)

// Post uploads the given data to Gateway3 and returns the corresponding CID.
func (c *Client) Post(data []byte) (string, error) {
	url, err := c.AuthPost(len(data))
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
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

type redirect struct {
	URL string `json:"url"`
}

// AuthPost gets the authorized URL from the Gateway3.
func (c *Client) AuthPost(size int) (string, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"%s/ipfs/?size=%d",
			c.endPoint,
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
