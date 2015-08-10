package goquadriga

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

func NewClient(id string, key string, secret string) *Client {
	return &Client{
		RootUrl:   V2URL,
		ClientId:  id,
		ApiKey:    key,
		ApiSecret: secret,
	}
}

func (c *Client) URL(res string) string {
	return fmt.Sprintf("%s%s", c.RootUrl, res)
}

func (c *Client) GetCurrentTradingInfo() (CurrentTrade, error) {
	var current CurrentTrade

	body, err := c.get(c.URL("ticker"))
	if err != nil {
		return current, err
	}
	err = json.Unmarshal(body, &current)
	return current, err
}

func (c *Client) GetOrderBook() (OrderBook, error) {
	var orders OrderBook

	body, err := c.get(c.URL("order_book"))
	if err != nil {
		return orders, err
	}
	err = json.Unmarshal(body, &orders)
	return orders, err
}

func (c *Client) GetTransactions() (TransactionResponse, error) {
	var transactions TransactionResponse

	body, err := c.get(c.URL("transactions"))
	if err != nil {
		return transactions, err
	}
	err = json.Unmarshal(body, &transactions)
	fmt.Printf("Results: %v\n", transactions)
	return transactions, err
}

func (c *Client) PostAccountBalance() (AccountBalance, error) {
	var balance AccountBalance

	auth := c.makeSig()
	payload, err := json.Marshal(auth)
	if err != nil {
		return balance, err
	}
	fmt.Println(string(payload))

	body, err := c.post(c.URL("balance"), payload)
	if err != nil {
		return balance, err
	}
	err = json.Unmarshal(body, &balance)
	return balance, nil
}

func (c *Client) get(url string) ([]byte, error) {
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

func (c *Client) post(url string, payload []byte) ([]byte, error) {
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
