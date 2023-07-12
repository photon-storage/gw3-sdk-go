package gw3

import (
	"fmt"
	"net/http"
)

// Pin requests Gateway3 to pin the given CID.
func (c *Client) Pin(cid string) error {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"%s/api/v0/pin/add?arg=%s",
			c.endPoint,
			cid,
		),
		nil,
	)
	if err != nil {
		return err
	}

	return c.callGateway(req, nil)
}

// Unpin requests Gateway3 to unpin the given CID.
func (c *Client) Unpin(cid string) error {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(
			"%s/api/v0/pin/rm?arg=%s",
			c.endPoint,
			cid,
		),
		nil,
	)
	if err != nil {
		return err
	}

	return c.callGateway(req, nil)
}
