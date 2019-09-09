package rpc

type ListBlocksTransactionsCommand struct {
	FirstBlock int64
	LastBlock int64
}

func (ListBlocksTransactionsCommand) Method() string {
	return "omni_listblockstransactions"
}

func (ListBlocksTransactionsCommand) ID() string {
	return "1"
}

func (cmd ListBlocksTransactionsCommand) Params() []interface{} {
	return []interface{}{cmd.FirstBlock, cmd.LastBlock}
}


