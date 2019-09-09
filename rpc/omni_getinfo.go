package rpc

type GetInfoCommand struct {}

func (GetInfoCommand) Method() string {
	return "omni_getinfo"
}

func (GetInfoCommand) ID() int {
	return 1
}

func (GetInfoCommand) Params() []interface{} {
	return []interface{}{}
}