package infura

type Block struct {
	Author           string           `json:"author"`
	Difficulty       string           `json:"difficulty"`
	ExtraData        string           `json:"extra_data"`
	GasLimit         string           `json:"gas_limit"`
	GasUsed          string           `json:"gas_used"`
	Hash             string           `json:"hash"`
	LogsBloom        string           `json:"logs_bloom"`
	Miner            string           `json:"miner"`
	MixHash          string           `json:"mix_hash"`
	Nonce            string           `json:"nonce"`
	Number           string           `json:"number"`
	ParentHash       string           `json:"parent_hash"`
	ReceiptsRoot     string           `json:"receipts_root"`
	SealFields       []string         `json:"seal_fields"`
	SHA3Uncles       string           `json:"sha3_uncles"`
	Size             string           `json:"size"`
	StateRoot        string           `json:"state_root"`
	Timestamp        string           `json:"timestamp"`
	TotalDifficulty  string           `json:"total_difficulty"`
	Transactions     []RawTransaction `json:"transactions"`
	TransactionsRoot string           `json:"transactions_root"`
	Uncles           []string         `json:"uncles"`
}
