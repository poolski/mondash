package main

import (
	"encoding/json"
	"fmt"
	"os"
	//"io/ioutil"
	"net/http"
	"time"
)

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

type Config struct {
	accountID string
	token     string
}

type Transaction struct {
	ID          string      `json:"id"`
	Created     time.Time   `json:"created"`
	Description string      `json:"description"`
	Amount      int         `json:"amount"`
	Currency    string      `json:"currency"`
	Merchant    interface{} `json:"merchant"`
	Notes       string      `json:"notes"`
	Metadata    struct {
	} `json:"metadata"`
	AccountBalance int           `json:"account_balance"`
	Attachments    []interface{} `json:"attachments"`
	Category       string        `json:"category"`
	IsLoad         bool          `json:"is_load"`
	Settled        time.Time     `json:"settled"`
	LocalAmount    int           `json:"local_amount"`
	LocalCurrency  string        `json:"local_currency"`
	Updated        time.Time     `json:"updated"`
	AccountID      string        `json:"account_id"`
	Counterparty   struct {
	} `json:"counterparty"`
	Scheme     string `json:"scheme"`
	DedupeID   string `json:"dedupe_id"`
	Originator bool   `json:"originator"`
}

func ReadConfig() Config {
	file, err := os.Open("conf.json")
	if err != nil {
		fmt.Print(err)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(config.accountID)
	return config
}

func main() {
	var config = ReadConfig()

	var accountID = config.accountID
	var token = config.token

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.getmondo.co.uk/transactions", nil)

	q := req.URL.Query()
	q.Add("account_id", accountID)
	req.Header.Add("Authorization", "Bearer "+token)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	decoder := json.NewDecoder(res.Body)
	var transactions Transactions
	err = decoder.Decode(&transactions)

	if err != nil {
		fmt.Print(err)
	}
	for _, t := range transactions.Transactions {
		fmt.Println(t)
	}

}
