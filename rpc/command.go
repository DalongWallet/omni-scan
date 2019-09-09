package rpc

import (
	"encoding/json"
	"fmt"
)

type command interface {
	ID()     int
	Method() string
	Params() []interface{}
}

func marshalCmd(cmd command) ([]byte, error) {
	return json.Marshal(newRpcRequest(cmd))
}

type rpcReqeust struct {
	ID      int            `json:"id"`
	JsonRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []interface{}     `json:"params"`
}

func newRpcRequest(cmd command) (*rpcReqeust) {
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

func (resp *rpcResponse) result() ([]byte, error) {
	if resp.Error != nil {
		return []byte{}, resp.Error
	}
	return json.Marshal(resp.Result)
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *rpcError) Error() string {
	return fmt.Sprintf("errCode: %d, errMsg: %s", e.Code, e.Message)
}
