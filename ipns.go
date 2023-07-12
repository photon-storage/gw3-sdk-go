package gw3

import (
	"fmt"
	"net/http"
)

type ipnsResp struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CreateIPNS creates a new IPNS record and binds it to the given CID.
// This function should only be used for creating a new record.
// To update an existing IPNS record, use the UpdateIPNS interface.
func (c *Client) CreateIPNS(value string) (string, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v0/name/create", c.endPoint),
		nil,
	)
	if err != nil {
		return "", err
	}

	query := req.URL.Query()
	query.Set("arg", value)
	req.URL.RawQuery = query.Encode()

	r := &ipnsResp{}
	return r.Name, c.callGateway(req, r)
}

// UpdateIPNS updates the value for the IPNS record specified by the given name.
func (c *Client) UpdateIPNS(name, value string) error {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v0/name/publish", c.endPoint),
		nil,
	)
	if err != nil {
		return err
	}

	query := req.URL.Query()
	query.Set("key", name)
	query.Set("arg", value)
	req.URL.RawQuery = query.Encode()

	r := &ipnsResp{}
	return c.callGateway(req, r)
}
