package rpc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

var DefaultRpcClient = &RpcClient{
	Host:   "127.0.0.1",
	Port:   "8332",
	Client: &http.Client{},
}

type RpcClient struct {
	Host string
	Port string
	*http.Client
}

func (client *RpcClient) SendJsonRpc(method string, params ...interface{}) ([]byte, error) {
	url := fmt.Sprintf("http://%s:%s", client.Host, client.Port)
	req := CommonRpcReq{
		// TODO: 取消 id 硬编码
		Id:         1,
		RpcVersion: "2.0",
		Method:     method,
		Params:     params,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "Marshal req failed")
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, errors.Wrap(err, "NewReq failed")
	}
	request.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("togreat:cd32d5e86e")))

	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "SendReq failed")
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Read data from resp.Body failed")
	}
	commonResp := &CommonRpcResp{}
	err = json.Unmarshal(respData, commonResp)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal data to CommonRpcResp failed")
	}

	return commonResp.result()
}

type CommonRpcReq struct {
	Id         int           `json:"id"`
	RpcVersion string        `json:"jsonrpc"`
	Method     string        `json:"method"`
	Params     []interface{} `json:"params"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *rpcError) Error() string {
	return fmt.Sprintf("errCode: %d, errMsg: %s", e.Code, e.Message)
}

type CommonRpcResp struct {
	Result json.RawMessage `json:"result"`
	Error  *rpcError       `json:"error"`
}

func (resp CommonRpcResp) result() ([]byte, error) {
	if resp.Error != nil {
		return []byte{}, errors.Wrap(resp.Error,"unexpect jsonRpc Resp")
	}
	return json.Marshal(resp.Result)
}
