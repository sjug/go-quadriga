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

type ErrorResponse struct {
	Error ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
