package omnicore

import "omni-scan/rpc"

type Client struct {
	RpcClient  		*rpc.OmniClient
	// TODO TxChan
}
