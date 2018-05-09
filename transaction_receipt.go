package infura

type TransactionReceipt struct {
	BlockHash         *string     `json:"blockHash"`
	BlockNumber       *string     `json:"blockNumber"`
	ContractAddress   *string     `json:"contractAddress"`
	CumulativeGasUsed *string     `json:"cumulativeGasUsed"`
	From              string      `json:"from"`
	GasUsed           string      `json:"gasUsed"`
	Logs              interface{} `json:"logs"`
	LogsBloom         string      `json:"logsBloom"`
	To                *string     `json:"to"`
	TransactionHash   string      `json:"transactionHash"`
	TransactionIndex  *string     `json:"transactionIndex"`
	Root              *string     `json:"root"`
	Status            string      `json:"status"`
}
