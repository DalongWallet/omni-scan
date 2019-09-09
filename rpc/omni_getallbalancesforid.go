package rpc

type GetAllbalancesForIdCommand struct {
	PropertyId int
}

func (GetAllbalancesForIdCommand) Method() string {
	return "omni_getallbalancesforid"
}

func (GetAllbalancesForIdCommand) ID() string {
	return "1"
}

func (cmd GetAllbalancesForIdCommand) Params() []interface{} {
	return []interface{}{cmd.PropertyId}
}


