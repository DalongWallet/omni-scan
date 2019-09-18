package rpc

type DecodeRawTransactionCommand struct {
	RawTx 			string
}

func (DecodeRawTransactionCommand) ID() int {
	return 1
}

func (DecodeRawTransactionCommand) Method() string {
	return "omni_decodetransaction"
}

func (cmd DecodeRawTransactionCommand) Params() []interface{} {
	return []interface{}{
		cmd.RawTx,
	}
}
