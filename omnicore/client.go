package omnicore

import "omni-scan/rpc"

type Client struct {
	RpcClient  		*rpc.OmniClient
	// TODO Local cmd
	// TODO TxChan
}
