package goquadriga

type BaseAuth struct {
	ApiKey    string `json:"key"`
	Signature string `json:"signature"`
	Nonce     string `json:"nonce"`
}

type AccountBalance struct {
	CadBalance   string `json:"cad_balance"`
	BtcBalance   string `json:"btc_balance"`
	CadReserved  string `json:"cad_reserved"`
	BtcReserved  string `json:"btc_reserved"`
	CadAvailable string `json:"cad_available"`
	BtcAvailable string `json:"btc_available"`
	Fee          string `json:"fee"`
}
