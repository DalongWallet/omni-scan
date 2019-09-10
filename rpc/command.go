package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type command interface {
	ID() int
	Method() string
	Params() []interface{}
}

func marshalCmd(cmd command) ([]byte, error) {
	return json.Marshal(newRpcRequest(cmd))
}

type rpcReqeust struct {
	ID      int           `json:"id"`
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

func newRpcRequest(cmd command) *rpcReqeust {
	return &rpcReqeust{
		ID:      cmd.ID(),
		JsonRPC: "2.0",
		Method:  cmd.Method(),
		Params:  cmd.Params(),
	}
}

type rpcResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *rpcError       `json:"error"`
}

func (resp *rpcResponse) result() (result []byte, err error) {
	if resp.Error != nil {
		return []byte{}, errors.Wrap(resp.Error, "Rpc BadResponse")
	}
	result, err = json.Marshal(resp.Result)
	err = errors.Wrap(err, "Marshal Rpc Response.Result failed")
	return
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *rpcError) Error() string {
	return fmt.Sprintf("errCode: %d, errMsg: %s", e.Code, e.Message)
}
