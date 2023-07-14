package gw3

import (
	"bytes"
	"context"
	"fmt"
	gohttp "net/http"
	"path"

	ipfsfiles "github.com/ipfs/boxo/files"

	"github.com/photon-storage/go-gw3/common/http"
	car "github.com/photon-storage/go-ipfs-car"
)

const (
	EmptyDAGRoot = "QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn"
)

// DAGAdd adds a new CID and path to the existing dag, generating a new dag root.
func (c *Client) DAGAdd(root, path string, data []byte) (string, error) {
	url, err := c.AuthDAGAdd(root, path, len(data))
	if err != nil {
		return "", err
	}

	req, err := gohttp.NewRequest(gohttp.MethodPut, url, bytes.NewReader(data))
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

// AuthDAGAdd requests Gateway3 for an authorized redirect URL for subsequently adding a new CID and path to the existing dag.
func (c *Client) AuthDAGAdd(root, filePath string, size int) (string, error) {
	p := path.Join(root, filePath)
	req, err := gohttp.NewRequest(
		gohttp.MethodPut,
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

// AuthDAGRemove requests Gateway3 for an authorized redirect URL for subsequently removing a path from the existing dag, generating a new dag root.
func (c *Client) AuthDAGRemove(root, filePath string) (string, error) {
	p := path.Join(root, filePath)
	req, err := gohttp.NewRequest(
		gohttp.MethodDelete,
		fmt.Sprintf("%s/ipfs/%s", c.endPoint, p),
		nil,
	)
	if err != nil {
		return "", err
	}

	var r redirect
	return r.URL, c.callGateway(req, &r)
}

// DAGRemove removes a path from the existing dag, generating a new dag root.
func (c *Client) DAGRemove(root, path string) (string, error) {
	url, err := c.AuthDAGRemove(root, path)
	if err != nil {
		return "", err
	}

	req, err := gohttp.NewRequest(gohttp.MethodDelete, url, nil)
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

// AuthDAGImport requests Gateway3 for an authorized redirect URL for uploading a CAR file.
func (c *Client) AuthDAGImport(size int, boundary string) (string, error) {
	req, err := gohttp.NewRequest(
		gohttp.MethodPost,
		fmt.Sprintf("%s/api/v0/dag/import", c.endPoint),
		nil,
	)
	if err != nil {
		return "", err
	}

	query := req.URL.Query()
	query.Set(http.ParamP3Size, fmt.Sprintf("%v", size))
	query.Set(http.ParamP3Boundary, boundary)
	req.URL.RawQuery = query.Encode()

	var r redirect
	return r.URL, c.callGateway(req, &r)
}

// DAGImport imports the given src input as a CAR format and returns its root CID. The `src` can be a path to a directory, a byte array or a io.Reader.
func (c *Client) DAGImport(src any) (string, error) {
	b := car.NewBuilder()
	v1car, err := b.Buildv1(context.TODO(), src, car.ImportOpts.CIDv1())
	if err != nil {
		return "", err
	}

	buf := bytes.Buffer{}
	if err := v1car.Write(&buf); err != nil {
		return "", err
	}

	r := ipfsfiles.NewMultiFileReader(
		ipfsfiles.NewMapDirectory(map[string]ipfsfiles.Node{
			"path": ipfsfiles.NewBytesFile(buf.Bytes()),
		}),
		true,
	)

	url, err := c.AuthDAGImport(buf.Len(), r.Boundary())
	if err != nil {
		return "", err
	}

	req, err := gohttp.NewRequest(gohttp.MethodPost, url, r)
	if err != nil {
		return "", err
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != gohttp.StatusOK {
		return "", fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	return v1car.Root().String(), nil
}
