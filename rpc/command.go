package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type command interface {
	ID() string
	Method() string
	Params() []interface{}
}

func marshalCmd(cmd command) ([]byte, error) {
	rawCmd, err := newRpcRequest(cmd)
	if err != nil {
		return nil, err
	}

	return json.Marshal(rawCmd)
}

type rpcReqeust struct {
	ID      string            `json:"id"`
	JsonRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
}

func newRpcRequest(cmd command) (*rpcReqeust, error) {
	params := cmd.Params()
	rawParams := make([]json.RawMessage, len(params))

	for i := range params {
		msg, err := json.Marshal(params[i])
		if err != nil {
			return nil, err
		}

		rawParams[i] = json.RawMessage(msg)
	}

	return &rpcReqeust{
		ID:      cmd.ID(),
		JsonRPC: "2.0",
		Method:  cmd.Method(),
		Params:  rawParams,
	}, nil
}

type rpcResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *rpcError       `json:"error"`
}

func newRpcResponse() *rpcResponse {
	return &rpcResponse{}
}

func (resp *rpcResponse) result() ([]byte, error) {
	if resp.Error != nil {
		return []byte{}, errors.Wrap(resp.Error, "unexpect jsonRpc Resp")
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
