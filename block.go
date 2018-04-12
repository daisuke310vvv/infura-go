package infura

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
)

type Block struct {
	Author           string   `json:"author"`
	Difficulty       int64    `json:"difficulty"`
	ExtraData        string   `json:"extra_data"`
	GasLimit         int      `json:"gas_limit"`
	GasUsed          int      `json:"gas_used"`
	Hash             string   `json:"hash"`
	LogsBloom        string   `json:"logs_bloom"`
	Miner            string   `json:"miner"`
	MixHash          string   `json:"mix_hash"`
	Nonce            *big.Int `json:"nonce"`
	Number           int      `json:"number"`
	ParentHash       string   `json:"parent_hash"`
	ReceiptsRoot     string   `json:"receipts_root"`
	SealFields       []string `json:"seal_fields"`
	SHA3Uncles       string   `json:"sha3_uncles"`
	Size             int      `json:"size"`
	StateRoot        string   `json:"state_root"`
	Timestamp        int      `json:"timestamp"`
	TotalDifficulty  *big.Int `json:"total_difficulty"`
	Transactions     []string `json:"transactions"`
	TransactionsRoot string   `json:"transactions_root"`
	Uncles           []string `json:"uncles"`
}

func (b *Block) UnmarshalJSON(data []byte) error {
	var blockMap map[string]interface{}
	err := json.Unmarshal(data, &blockMap)
	if err != nil {
		return err
	}

	for k, v := range blockMap {
		fmt.Printf("key: %v, value: %v\n", k, v)
		switch k {
		case "auther":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b.Author = c
		case "difficulty":
			c, ok := v.(string)
			if !ok {
				continue
			}
			d, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			b.Difficulty = d
		case "extra_data":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b.ExtraData = c
		case "gas_limit":
			c, ok := v.(string)
			if !ok {
				continue
			}
			g, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			b.GasLimit = int(g)
		case "gas_used":
			c, ok := v.(string)
			if !ok {
				continue
			}
			g, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			b.GasUsed = int(g)
		case "hash":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b.Hash = c
		case "logs_bloom":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b.LogsBloom = c
		case "miner":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b.Miner = c
		case "mix_hash":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b.MixHash = c
		case "nonce":
			c, ok := v.(string)
			if !ok {
				continue
			}
			s := Remove0x(c)
			bint, _ := big.NewInt(0).SetString(s, 16)
			b.Nonce = bint
		case "number":
			c, ok := v.(string)
			if !ok {
				continue
			}
			n, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			b.Number = int(n)
		case "parent_hash":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b.ParentHash = c
		case "receipts_root":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b.ReceiptsRoot = c
		case "seal_fields":
			c, ok := v.([]string)
			if !ok {
				continue
			}
			b.SealFields = c
		case "sha3_uncles":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b.SHA3Uncles = c
		case "size":
			c, ok := v.(string)
			if !ok {
				continue
			}
			s, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			b.Size = int(s)
		case "state_root":
			c, ok := v.(string)
			if !ok {
				continue
			}
			b.StateRoot = c
		case "timestamp":
			c, ok := v.(string)
			if !ok {
				continue
			}
			t, err := DecodeToInt(c)
			if err != nil {
				return err
			}
			b.Timestamp = int(t)
		case "total_difficulty":
			c, ok := v.(string)
			if !ok {
				continue
			}
			s := Remove0x(c)
			bint, ok := big.NewInt(0).SetString(s, 16)
			if !ok {
				b.TotalDifficulty = nil
			}
			b.TotalDifficulty = bint
		case "transactions":
			val := reflect.ValueOf(v)
			fmt.Println(val.Kind())
			if val.Kind() == reflect.Slice {
				for i := 0; i < val.Len(); i++ {
					b.Transactions = append(b.Transactions, fmt.Sprintf("%s", val.Index(i)))
				}
			}
		case "transactons_root":
			b.TransactionsRoot = v.(string)
		case "uncles":
			val := reflect.ValueOf(v)
			fmt.Println(val.Kind())
			if val.Kind() == reflect.Slice {
				for i := 0; i < val.Len(); i++ {
					b.Uncles = append(b.Uncles, fmt.Sprintf("%s", val.Index(i)))
				}
			}
		}
	}
	return nil
}
