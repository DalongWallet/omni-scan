package rpc

type ListPropertiesCommand struct {}

func (ListPropertiesCommand) Method() string {
	return "omni_listproperties"
}

func (ListPropertiesCommand) ID() int {
	return 1
}

func (cmd ListPropertiesCommand) Params() []interface{} {
	return []interface{}{}
}