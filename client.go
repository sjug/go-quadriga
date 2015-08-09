package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const V2URL = "https://api.quadrigacx.com/v2/"

type Client struct {
	RootUrl   string
	ClientId  string
	ApiKey    string
	ApiSecret string
}

type BaseAuth struct {
	ApiKey    string `json:"key"`
	Signature string `json:"signature"`
	Nonce     string `json:"nonce"`
}

func NewClient(id string, key string, secret string) *Client {
	return &Client{
		RootUrl:   V2URL,
		ClientId:  id,
		ApiKey:    key,
		ApiSecret: secret,
	}
}

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

type AccountBalance struct {
	CadBalance   string `json:"cad_balance"`
	BtcBalance   string `json:"btc_balance"`
	CadReserved  string `json:"cad_reserved"`
	BtcReserved  string `json:"btc_reserved"`
	CadAvailable string `json:"cad_available"`
	BtcAvailable string `json:"btc_available"`
	Fee          string `json:"fee"`
}

func main() {
	cl := NewClient("62915", "zzRDeTjONf", "si&ia0;:G0z2NMK5hpSFh]HQG[wDKFne+G1cQ1ch2")
	trade, err := cl.getCurrentTradingInfo()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Last trade: ", trade.Last, "\nVolume: ", trade.Volume, "\nHigh: ", trade.High, "\nLow: ", trade.Low)
	}

	orders, err := cl.getOrderBook()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Top sell price: ", orders.Asks[0][0], " Top sell amount: ", orders.Asks[0][1])
	}

	balance, err := cl.postAccountBalance()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Current CAD$ balance:", balance.CadBalance, " Current BTC balance:", balance.BtcBalance)
	}
}

func (c *Client) URL(res string) string {
	return fmt.Sprintf("%s%s", c.RootUrl, res)
}

func (c *Client) getCurrentTradingInfo() (CurrentTrade, error) {
	var current CurrentTrade

	body, err := c.Get(c.URL("ticker"))
	if err != nil {
		return current, err
	}
	err = json.Unmarshal(body, &current)
	return current, err
}

func (c *Client) getOrderBook() (OrderBook, error) {
	var orders OrderBook

	body, err := c.Get(c.URL("order_book"))
	if err != nil {
		return orders, err
	}
	err = json.Unmarshal(body, &orders)
	return orders, err
}

func (c *Client) postAccountBalance() (AccountBalance, error) {
	var balance AccountBalance

	auth := c.makeSig()
	payload, err := json.Marshal(auth)
	if err != nil {
		return balance, err
	}
	fmt.Println(string(payload))

	body, err := c.Post(c.URL("balance"), payload)
	if err != nil {
		return balance, err
	}
	err = json.Unmarshal(body, &balance)
	return balance, nil
}

func (c *Client) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) Post(url string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response status: ", resp.Status)
	fmt.Println("Response headers: ", resp.Header)
	fmt.Println("Response body: ", string(body))
	return body, nil
}

func (c *Client) makeSig() *BaseAuth {
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	secretHash := md5.New()
	secretHash.Write([]byte(c.ApiSecret))
	key := hex.EncodeToString(secretHash.Sum(nil))
	fmt.Println("The secret key is ", key)

	message := strings.Join([]string{timestamp, c.ClientId, c.ApiKey}, "")
	fmt.Println("The message is ", message)

	sig := hmac.New(sha256.New, []byte(key))
	sig.Write([]byte(message))

	base := &BaseAuth{ApiKey: c.ApiKey, Signature: hex.EncodeToString(sig.Sum(nil)), Nonce: timestamp}
	return base
}
