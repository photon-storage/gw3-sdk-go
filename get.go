package gw3

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Get retrieves data from the Gateway3 for the given CID.
func (c *Client) Get(cid string) ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/ipfs/%s",
			c.endPoint,
			cid,
		),
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err := c.sign(req); err != nil {
		return nil, err
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "call gateway API")
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// GetIPNS retrieves data from the IPFS network using the given IPNS.
func (c *Client) GetIPNS(name string) ([]byte, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/ipns/%s",
			c.endPoint,
			name,
		),
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err := c.sign(req); err != nil {
		return nil, err
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "call gateway API")
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
