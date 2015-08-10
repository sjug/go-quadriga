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

type OrderId struct {
	ID string `json:"id"`
}

type LookupOrderResponse struct {
	Amount  string `json:"amount"`
	Book    string `json:"book"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	ID      string `json:"id"`
	Price   string `json:"price"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}

type TransactionResponse []Transaction

type Transaction struct {
	Amount string `json:"amount"`
	Date   string `json:"date"`
	Price  string `json:"price"`
	Tid    uint   `json:"tid"`
}

type OpenOrdersResponse []OpenOrder

type OpenOrder struct {
	Amount   string `json:"amount"`
	Datetime string `json:"datetime"`
	ID       string `json:"id"`
	Price    string `json:"price"`
	Status   string `json:"status"`
	Type     string `json:"type"`
}
