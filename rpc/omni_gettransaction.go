package rpc

type GetTransactionCommand struct {
	TxId string
}

func (GetTransactionCommand) Method() string {
	return "omni_gettransaction"
}

func (GetTransactionCommand) ID() int {
	return 1
}

func (cmd GetTransactionCommand) Params() []interface{} {
	return []interface{}{cmd.TxId}
}

