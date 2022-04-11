package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
// 	"strconv"
)

type TxResponse struct {
	Jsonrpc string            `json:"jsonrpc"`
	ID      int               `json:"id"`
	Result  *TxResponseResult `json:"result"`
}

type TxResponseResult struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	Type             string `json:"type"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

func main() {
	url := "http://localhost:8545"
	for i := len(AllStateSyncs) - 1; i >= 0; i -= 1 {

		var jsonStr = []byte(`{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["` + AllStateSyncs[i] + `"],"id":0}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		var result TxResponse
		if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Bor: Cannot unmarshal JSON")
			fmt.Println(string(body))
		}

		if result.Result == nil {
			fmt.Println("Missing :  " + AllStateSyncs[i])
			break
		} else {
// 			n, err := strconv.ParseUint(result.Result.BlockNumber[2:], 16, 64)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println(n)
			if i == 0 {
				fmt.Println("No Missing StateSyncs")
			}
		}
	}

}
