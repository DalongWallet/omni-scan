package rpc


type GetBalanceCommand struct {
	Address string
	PropertyId int
}

func (GetBalanceCommand) Method() string {
	return "omni_getbalance"
}

func (GetBalanceCommand) ID() string {
	return "1"
}

func (cmd GetBalanceCommand) Params() []interface{} {
	return []interface{}{cmd.Address, cmd.PropertyId}
}


