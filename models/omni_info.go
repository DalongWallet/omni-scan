package models


type Alert struct {
	AlertType    string `json:"alerttype"`
	AlertExpiry  string `json:"alertexpiry"`
	AlertMessage string `json:"alertmessage"`
}

type OmniInfoResult struct {
	OmniCoreVersionInt      int     `json:"omnicoreversion_int"`
	OmniCoreVersion         string  `json:"omnicoreversion"`
	MasterCoreVersion       string  `json:"mastercoreversion"`
	BitcoinCoreVersion      string  `json:"bitcoincoreversion"`
	CommitInfo              string  `json:"commitinfo"`
	BlockHeight             int64     `json:"block"`
	BlockTime               int64   `json:"blocktime"`
	LatestBlockTransactions int64   `json:"blocktransactions"`
	TotalTransactions       int64   `json:"totaltransactions"`
	Alerts                  []Alert `json:"alerts"`
}