package rpc

type ListBlockTransactionsCommand struct {
	Index int
}

func (ListBlockTransactionsCommand) Method() string {
	return "omni_listblocktransactions"
}

func (ListBlockTransactionsCommand) ID() int {
	return 1
}

func (cmd ListBlockTransactionsCommand) Params() []interface{} {
	return []interface{}{cmd.Index}
}
