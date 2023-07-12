package gw3

import (
	"encoding/base64"
	"net/http"

	"github.com/pkg/errors"
)

const defaultEndPoint = "https://gw3.io"

// Client represents a Gateway3 client for interacting with the Gateway3.
type Client struct {
	endPoint string
	key      string
	secret   []byte

	hc *http.Client
}

// NewClient creates a new Gateway3 client with the provided access key and
// access secret.
func NewClient(accessKey, accessSecret string) (*Client, error) {
	secret, err := base64.URLEncoding.DecodeString(accessSecret)
	if err != nil {
		return nil, errors.WithMessage(err, "decode secret")
	}

	return &Client{
		endPoint: defaultEndPoint,
		key:      accessKey,
		secret:   secret,
		hc:       &http.Client{},
	}, nil
}
