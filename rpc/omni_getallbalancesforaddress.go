package rpc

type GetAllBalancesForAddressCommand struct {
	Address string
}

func (GetAllBalancesForAddressCommand) Method() string {
	return "omni_getallbalancesforaddress"
}

func (GetAllBalancesForAddressCommand) ID() string {
	return "1"
}

func (cmd GetAllBalancesForAddressCommand) Params() []interface{} {
	return []interface{}{cmd.Address}
}

