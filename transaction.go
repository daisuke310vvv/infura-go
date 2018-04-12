package infura

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Transaction struct {
	From     string   `json:"from"`
	Gas      int      `json:"gas"`
	GasPrice *big.Int `json:"gas_price"`
	To       *string  `json:"to"`
	Value    *big.Int `json:"value"`
	Data     string   `json:"data"`
}

func (tx *Transaction) ToMap() map[string]interface{} {
	txMap := make(map[string]interface{})
	txMap["from"] = tx.From
	txMap["gas"] = "0x" + ToHexString(tx.Gas)

	if tx.GasPrice != nil {
		txMap["gasPrice"] = "0x" + tx.GasPrice.Text(16)
	}

	if tx.To != nil {
		txMap["to"] = *tx.To
	}

	if tx.Value != nil {
		txMap["value"] = "0x" + tx.Value.Text(16)
	}

	if len(tx.Data) == 0 {
		txMap["data"] = "0x"
	} else {
		txMap["data"] = tx.Data
	}

	return txMap
}

type RawTransaction struct {
	Transaction
	BlockHash        *string `json:"block_hash"`
	BlockNumber      *int    `json:"block_number"`
	Hash             string  `json:"hash"`
	Input            string  `json:"input"`
	Nonce            int     `json:"nonce"`
	R                string  `json:"r"`
	S                string  `json:"s"`
	TransactionIndex *int    `json:"transaction_index"`
	V                int     `json:"v"`

	// Parity only
	// Condition *string `json:"condition"`
	// ChainId   *int    `json:"chain_id"`
	// Creates   *string `json:"creates"`
	// PublicKey *string `json:"public_key"`
	// Raw       *string `json:"raw"`
	// StandardV *int    `json:"standard_v"`
}

func (t *RawTransaction) UnmarshalJSON(data []byte) error {
	var txMap map[string]interface{}
	err := json.Unmarshal(data, &txMap)
	if err != nil {
		return err
	}

	for k, v := range txMap {
		fmt.Printf("key: %v, value: %v\n", k, v)
		switch k {
		case "from":
			c, ok := v.(string)
			if !ok {
				continue
			}
			t.From = c
		case "gas":
			c, ok := v.(string)
			if !ok {
				continue
			}
			g, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			t.Gas = int(g)
		case "gas_price":
			c, ok := v.(string)
			if !ok {
				continue
			}
			s := Remove0x(c)
			bint, _ := big.NewInt(0).SetString(s, 16)
			t.GasPrice = bint
		case "to":
			c, ok := v.(string)
			if !ok {
				continue
			}
			t.To = &c
		case "value":
			c, ok := v.(string)
			if !ok {
				continue
			}
			s := Remove0x(c)
			bint, _ := big.NewInt(0).SetString(s, 16)
			t.Value = bint
		case "data":
			c, ok := v.(string)
			if !ok {
				continue
			}
			t.Data = c
		case "block_hash":
			c, ok := v.(string)
			if !ok {
				continue
			}
			blockHash := c
			t.BlockHash = &blockHash
		case "block_number":
			c, ok := v.(string)
			if !ok {
				continue
			}
			n, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			n64 := int(n)
			t.BlockNumber = &n64
		case "hash":
			c, ok := v.(string)
			if !ok {
				continue
			}
			t.Hash = c
		case "input":
			c, ok := v.(string)
			if !ok {
				continue
			}
			t.Input = c
		case "nonce":
			c, ok := v.(string)
			if !ok {
				continue
			}
			g, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			t.Nonce = int(g)
		case "r":
			c, ok := v.(string)
			if !ok {
				continue
			}
			t.R = c
		case "s":
			c, ok := v.(string)
			if !ok {
				continue
			}
			t.S = c
		case "v":
			c, ok := v.(string)
			if !ok {
				continue
			}
			g, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			t.V = int(g)
		case "transaction_index":
			c, ok := v.(string)
			if !ok {
				continue
			}
			n, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			n64 := int(n)
			t.TransactionIndex = &n64
		}
	}
	return nil
}
