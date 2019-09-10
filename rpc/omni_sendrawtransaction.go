package rpc

type SendRawTransactionCommand struct {
	FromAddress		string
	Hex   			string
}

func (SendRawTransactionCommand) ID() int {
	return 1
}

func (SendRawTransactionCommand) Method() string {
	return "omni_sendrawtx"
}

func (cmd SendRawTransactionCommand) Params() []interface{} {
	return []interface{}{
		cmd.FromAddress,
		cmd.Hex,
	}
}
