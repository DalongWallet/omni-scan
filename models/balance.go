package models

type TokenBalance struct {
	Balance string `json:"balance"`
	Reserved string `json:"reserved"`
	Frozen string `json:"frozen"`
}

type AddressTokenBalance struct {
	Address string  `json:"address"`
	TokenBalance
}

type PropertyTokenBalance struct {
	PropertyId int `json:"propertyid"`
	Name string `json:"name"`
	TokenBalance
}
