package infura

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type TransactionReceipt struct {
	BlockHash         *string     `json:"blockHash"`
	BlockNumber       *int        `json:"blockNumber"`
	ContractAddress   *string     `json:"contractAddress"`
	CumulativeGasUsed *big.Int    `json:"cumulativeGasUsed"`
	From              string      `json:"from"`
	GasUsed           int         `json:"gasUsed"`
	Logs              interface{} `json:"logs"`
	LogsBloom         string      `json:"logsBloom"`
	To                *string     `json:"to"`
	TransactionHash   string      `json:"transactionHash"`
	TransactionIndex  *int        `json:"transactionIndex"`
	Root              *string     `json:"root"`
	Status            bool        `json:"status"`
}

func (r *TransactionReceipt) UnmarshalJSON(data []byte) error {
	var receiptMap map[string]interface{}
	err := json.Unmarshal(data, &receiptMap)
	if err != nil {
		return err
	}

	for k, v := range receiptMap {
		fmt.Printf("key: %v, value: %v\n", k, v)
		switch k {
		case "blockHash":
			c, ok := v.(string)
			if !ok {
				continue
			}
			r.BlockHash = &c
		case "blockNumber":
			c, ok := v.(string)
			if !ok {
				continue
			}
			n, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			n64 := int(n)
			r.BlockNumber = &n64
		case "contractAddress":
			c, ok := v.(string)
			if !ok {
				continue
			}
			r.ContractAddress = &c
		case "cumulativeGasUsed":
			c, ok := v.(string)
			if !ok {
				continue
			}
			s := Remove0x(c)
			bint, _ := big.NewInt(0).SetString(s, 16)
			r.CumulativeGasUsed = bint
		case "from":
			c, ok := v.(string)
			if !ok {
				continue
			}
			r.From = c
		case "gasUsed":
			c, ok := v.(string)
			if !ok {
				continue
			}
			g, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			r.GasUsed = int(g)
		case "logs":
			r.Logs = v
		case "logsBloom":
			c, ok := v.(string)
			if !ok {
				continue
			}
			r.LogsBloom = c
		case "to":
			c, ok := v.(string)
			if !ok {
				continue
			}
			r.To = &c
		case "transactionHash":
			c, ok := v.(string)
			if !ok {
				continue
			}
			r.TransactionHash = c
		case "transactionIndex":
			c, ok := v.(string)
			if !ok {
				continue
			}
			n, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			n64 := int(n)
			r.TransactionIndex = &n64
		case "root":
			c, ok := v.(string)
			if !ok {
				continue
			}
			root := c
			r.Root = &root
		case "status":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b := Remove0x(c)
			r.Status = b == "1"
		}
	}
	return nil
}
