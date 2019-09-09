package rpc

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var DefaultOmniClient = &OmniClient{
	config: &ConnConfig{
		Host: "127.0.0.1:8332",
		User: os.Getenv("OMNICORE_USER"),
		Pwd:  os.Getenv("OMNICORE_PWD"),
	},
	httpClient: newHTTPClient(),
}

type ConnConfig struct {
	Host string
	User string
	Pwd  string
}

type OmniClient struct {
	config     *ConnConfig
	httpClient *http.Client
}

func (c *OmniClient) Exec(cmd command) ([]byte, error) {
	body, err := marshalCmd(cmd)
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest(http.MethodPost, "http://"+c.config.Host, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "New request failed")
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.config.User, c.config.Pwd)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "SendReq failed")
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Read data from resp.Body failed")
	}

	rpcResp := newRpcResponse()
	err = json.Unmarshal(respBytes, &rpcResp)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal data to CommonRpcResp failed")
	}

	return rpcResp.result()
}

func NewOmniClient(config *ConnConfig) (*OmniClient) {
	return &OmniClient{
		config:     config,
		httpClient: newHTTPClient(),
	}
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 4 * time.Second,
			IdleConnTimeout:       5 * 60 * time.Second,
		},
	}
}
