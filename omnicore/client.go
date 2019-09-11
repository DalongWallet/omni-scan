package omnicore

import "github.com/DalongWallet/omni-scan/rpc"

type Client struct {
	RpcClient  		*rpc.OmniClient
	// TODO TxChan
}
