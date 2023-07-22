package gw3

import (
	"bytes"
	"encoding/json"
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

// ImportIPNSReq defines the request payload for importing an IPNS record
type ImportIPNSReq struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	SecretKey string `json:"secret_key"`
	Format    string `json:"format"`
	Seq       uint64 `json:"seq"`
}

// ImportIPNS imports an IPNS record using a user-side generated private key.
func (c *Client) ImportIPNS(body *ImportIPNSReq) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v0/name/import", c.endPoint),
		bytes.NewReader(b),
	)
	if err != nil {
		return err
	}

	return c.callGateway(req, nil)
}
