package goquadriga

type CurrentTrade struct {
	Ask       string `json:"ask"`
	Bid       string `json:"bid"`
	High      string `json:"high"`
	Last      string `json:"last"`
	Low       string `json:"low"`
	Timestamp string `json:"timestamp"`
	Volume    string `json:"volume"`
	Vwap      string `json:"vwap"`
}

type OrderBook struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Timestamp string     `json:"timestamp"`
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
