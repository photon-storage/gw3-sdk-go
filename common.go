package gw3

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func (c *Client) sign(r *http.Request) error {
	query := r.URL.Query()
	query.Set("ts", fmt.Sprintf("%v", time.Now().Unix()))
	r.URL.RawQuery = query.Encode()

	data := fmt.Sprintf("%s\n%s\n%s", r.Method, r.URL.Path, r.URL.Query().Encode())
	mac := hmac.New(sha256.New, c.secret)
	if _, err := mac.Write([]byte(data)); err != nil {
		return err
	}

	sign := base64.URLEncoding.EncodeToString(mac.Sum(nil))
	r.Header.Set("X-Access-Key", c.key)
	r.Header.Set("X-Access-Signature", sign)
	return nil
}

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func (c *Client) callGateway(req *http.Request, data any) error {
	if err := c.sign(req); err != nil {
		return err
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return errors.WithMessage(err, "call gateway API")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	r := &response{Data: data}
	if err := json.Unmarshal(body, r); err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(r.Msg)
	}

	return nil
}
