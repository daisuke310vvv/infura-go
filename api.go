package infura

type input interface {
	method() string
	params() interface{}
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

type Web3ClientVersionInput struct {
}

func (i *Web3ClientVersionInput) method() string {
	return "web3_clientVersion"
}

func (i *Web3ClientVersionInput) params() interface{} {
	return []string{}
}

type Web3ClientVersionOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) Web3ClientVersionRequest(input *Web3ClientVersionInput) (req *Request, output *Web3ClientVersionOutput) {
	if input != nil {
		input = &Web3ClientVersionInput{}
	}

	output = &Web3ClientVersionOutput{}
	req = i.newRequest(input, output)
	return
}

type Web3Sha3Input struct {
	Value string
}

func (i *Web3Sha3Input) method() string {
	return "web3_sha3"
}

func (i *Web3Sha3Input) params() interface{} {
	return []string{i.Value}
}

type Web3Sha3Output struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) Web3Sha3Request(input *Web3Sha3Input) (req *Request, output *Web3Sha3Output) {
	if input == nil {
		input = &Web3Sha3Input{}
	}

	output = &Web3Sha3Output{}
	req = i.newRequest(input, output)
	return
}

type NetVersionInput struct {
}

func (i *NetVersionInput) method() string {
	return "net_version"
}

func (i *NetVersionInput) params() interface{} {
	return []string{}
}

type NetVersionOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) NetVersionRequest(input *NetVersionInput) (req *Request, output *NetVersionOutput) {
	if input == nil {
		input = &NetVersionInput{}
	}

	output = &NetVersionOutput{}
	req = i.newRequest(input, output)
	return
}

type NetListeningInput struct {
}

func (i *NetListeningInput) method() string {
	return "net_listening"
}

func (i *NetListeningInput) params() interface{} {
	return []string{}
}

type NetListeningOutput struct {
	Error  *Error `json:"error"`
	Result bool   `json:"result"`
}

func (i *Infura) NetListeningRequest(input *NetListeningInput) (req *Request, output *NetListeningOutput) {
	if input == nil {
		input = &NetListeningInput{}
	}

	output = &NetListeningOutput{}
	req = i.newRequest(input, output)
	return
}

type NetPeerCountInput struct {
}

func (i *NetPeerCountInput) method() string {
	return "net_peerCount"
}

func (i *NetPeerCountInput) params() interface{} {
	return []string{}
}

type NetPeerCountOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) NetPeerCountRequest(input *NetPeerCountInput) (req *Request, output *NetPeerCountOutput) {
	if input == nil {
		input = &NetPeerCountInput{}
	}

	output = &NetPeerCountOutput{}
	req = i.newRequest(input, output)
	return
}

type EthProtocolVersionInput struct {
}

func (i *EthProtocolVersionInput) method() string {
	return "eth_protocolVersion"
}

func (i *EthProtocolVersionInput) params() interface{} {
	return []string{}
}

type EthProtocolVersionOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthProtocolVersionRequest(input *EthProtocolVersionInput) (req *Request, output *EthProtocolVersionOutput) {
	if input == nil {
		input = &EthProtocolVersionInput{}
	}

	output = &EthProtocolVersionOutput{}
	req = i.newRequest(input, output)
	return
}

type EthSyncingInput struct {
}

func (i *EthSyncingInput) method() string {
	return "eth_syncing"
}

func (i *EthSyncingInput) params() interface{} {
	return []string{}
}

type EthSyncingOutput struct {
	Error *Error `json:"error"`
	// Result
	/*
		{
			"result": {
				startingBlock: '0x384',
				currentBlock: '0x386',
				highestBlock: '0x454',
			}
		}
		or
		{
			"result": false
		}
	*/
	Result interface{} `json:"result"`
}

func (i *Infura) EthSyncingRequest(input *EthSyncingInput) (req *Request, output *EthSyncingOutput) {
	if input == nil {
		input = &EthSyncingInput{}
	}

	output = &EthSyncingOutput{}
	req = i.newRequest(input, output)
	return
}

type EthMiningInput struct {
}

func (i *EthMiningInput) method() string {
	return "eth_mining"
}

func (i *EthMiningInput) params() interface{} {
	return []string{}
}

type EthMiningOutput struct {
	Error  *Error `json:"error"`
	Result bool   `json:"result"`
}

func (i *Infura) EthMiningRequest(input *EthMiningInput) (req *Request, output *EthMiningOutput) {
	if input == nil {
		input = &EthMiningInput{}
	}

	output = &EthMiningOutput{}
	req = i.newRequest(input, output)
	return
}

type EthHashrateInput struct {
}

func (i *EthHashrateInput) method() string {
	return "eth_hashrate"
}

func (i *EthHashrateInput) params() interface{} {
	return []string{}
}

type EthHashrateOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthHashrateRequest(input *EthHashrateInput) (req *Request, output *EthHashrateOutput) {
	if input != nil {
		input = &EthHashrateInput{}
	}

	output = &EthHashrateOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGasPriceInput struct {
}

func (i *EthGasPriceInput) method() string {
	return "eth_gasPrice"
}

func (i *EthGasPriceInput) params() interface{} {
	return []string{}
}

type EthGasPriceOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthGasPriceRequest(input *EthGasPriceInput) (req *Request, output *EthGasPriceOutput) {
	if input != nil {
		input = &EthGasPriceInput{}
	}

	output = &EthGasPriceOutput{}
	req = i.newRequest(input, output)
	return
}

type EthAccountsInput struct {
}

func (i *EthAccountsInput) method() string {
	return "eth_accounts"
}

func (i *EthAccountsInput) params() interface{} {
	return []string{}
}

type EthAccountsOutput struct {
	Error  *Error   `json:"error"`
	Result []string `json:"result"`
}

func (i *Infura) EthAccountsRequest(input *EthAccountsInput) (req *Request, output *EthAccountsOutput) {
	if input != nil {
		input = &EthAccountsInput{}
	}

	output = &EthAccountsOutput{}
	req = i.newRequest(input, output)
	return
}

type EthBlockNumberInput struct {
}

func (i *EthBlockNumberInput) method() string {
	return "eth_blockNumber"
}

func (i *EthBlockNumberInput) params() interface{} {
	return []string{}
}

type EthBlockNumberOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthBlockNumberRequest(input *EthBlockNumberInput) (req *Request, output *EthBlockNumberOutput) {
	if input != nil {
		input = &EthBlockNumberInput{}
	}

	output = &EthBlockNumberOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetBalanceInput struct {
	Address        string
	BlockParameter BlockParameter
}

func (i *EthGetBalanceInput) method() string {
	return "eth_getBalance"
}

func (i *EthGetBalanceInput) params() interface{} {
	return []string{i.Address, i.BlockParameter.Value}
}

type EthGetBalanceOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthGetBalanceRequest(input *EthGetBalanceInput) (req *Request, output *EthGetBalanceOutput) {
	if input == nil {
		input = &EthGetBalanceInput{}
	}

	output = &EthGetBalanceOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetTransactionCountInput struct {
	Address        string
	BlockParameter BlockParameter
}

func (i *EthGetTransactionCountInput) method() string {
	return "eth_getTransactionCount"
}

func (i *EthGetTransactionCountInput) params() interface{} {
	return []string{i.Address, i.BlockParameter.Value}
}

type EthGetTransactionCountOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthGetTransactionCountRequest(input *EthGetTransactionCountInput) (req *Request, output *EthGetTransactionCountOutput) {
	if input == nil {
		input = &EthGetTransactionCountInput{}
	}

	output = &EthGetTransactionCountOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetBlockTransactionCountByHashInput struct {
	Hash string
}

func (i *EthGetBlockTransactionCountByHashInput) method() string {
	return "eth_getBlockTransactionCountByHash"
}

func (i *EthGetBlockTransactionCountByHashInput) params() interface{} {
	return []string{i.Hash}
}

type EthGetBlockTransactionCountByHashOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthGetBlockTransactionCountByHashRequest(input *EthGetBlockTransactionCountByHashInput) (req *Request, output *EthGetBlockTransactionCountByHashOutput) {
	if input == nil {
		input = &EthGetBlockTransactionCountByHashInput{}
	}

	output = &EthGetBlockTransactionCountByHashOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetBlockTransactionCountByNumberInput struct {
	BlockParameter BlockParameter
}

func (i *EthGetBlockTransactionCountByNumberInput) method() string {
	return "eth_getBlockTransactionCountByNumber"
}

func (i *EthGetBlockTransactionCountByNumberInput) params() interface{} {
	return []string{i.BlockParameter.Value}
}

type EthGetBlockTransactionCountByNumberOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthGetBlockTransactionCountByNumberRequest(input *EthGetBlockTransactionCountByNumberInput) (req *Request, output *EthGetBlockTransactionCountByNumberOutput) {
	if input == nil {
		input = &EthGetBlockTransactionCountByNumberInput{}
	}

	output = &EthGetBlockTransactionCountByNumberOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetUncleCountByBlockHashInput struct {
	Hash string
}

func (i *EthGetUncleCountByBlockHashInput) method() string {
	return "eth_getUncleCountByBlockHash"
}

func (i *EthGetUncleCountByBlockHashInput) params() interface{} {
	return []string{i.Hash}
}

type EthGetUncleCountByBlockHashOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthGetUncleCountByBlockHashRequest(input *EthGetUncleCountByBlockHashInput) (req *Request, output *EthGetUncleCountByBlockHashOutput) {
	if input == nil {
		input = &EthGetUncleCountByBlockHashInput{}
	}

	output = &EthGetUncleCountByBlockHashOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetUncleCountByBlockNumberInput struct {
	BlockParameter BlockParameter
}

func (i *EthGetUncleCountByBlockNumberInput) method() string {
	return "eth_getUncleCountByBlockNumber"
}

func (i *EthGetUncleCountByBlockNumberInput) params() interface{} {
	return []string{i.BlockParameter.Value}
}

type EthGetUncleCountByBlockNumberOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthGetUncleCountByBlockNumberRequest(input *EthGetUncleCountByBlockNumberInput) (req *Request, output *EthGetUncleCountByBlockNumberOutput) {
	if input == nil {
		input = &EthGetUncleCountByBlockNumberInput{}
	}

	output = &EthGetUncleCountByBlockNumberOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetCodeInput struct {
	Address        string
	BlockParameter BlockParameter
}

func (i *EthGetCodeInput) method() string {
	return "eth_getCode"
}

func (i *EthGetCodeInput) params() interface{} {
	return []string{i.Address, i.BlockParameter.Value}
}

type EthGetCodeOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthGetCodeRequest(input *EthGetCodeInput) (req *Request, output *EthGetCodeOutput) {
	if input == nil {
		input = &EthGetCodeInput{}
	}

	output = &EthGetCodeOutput{}
	req = i.newRequest(input, output)
	return
}

type EthSendRawTransactionInput struct {
	Data string
}

func (i *EthSendRawTransactionInput) method() string {
	return "eth_sendRawTransaction"
}

func (i *EthSendRawTransactionInput) params() interface{} {
	return []string{i.Data}
}

type EthSendRawTransactionOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthSendRawTransactionRequest(input *EthSendRawTransactionInput) (req *Request, output *EthSendRawTransactionOutput) {
	if input == nil {
		input = &EthSendRawTransactionInput{}
	}

	output = &EthSendRawTransactionOutput{}
	req = i.newRequest(input, output)
	return
}

type EthCallInput struct {
	Transaction    Transaction
	BlockParameter BlockParameter
}

func (i *EthCallInput) method() string {
	return "eth_call"
}

func (i *EthCallInput) params() interface{} {
	return []interface{}{i.Transaction.toMap(), i.BlockParameter.Value}
}

type EthCallOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthCallRequest(input *EthCallInput) (req *Request, output *EthCallOutput) {
	if input == nil {
		input = &EthCallInput{}
	}

	output = &EthCallOutput{}
	req = i.newRequest(input, output)
	return
}

type EthEstimateGasInput struct {
	Transaction Transaction
}

func (i *EthEstimateGasInput) method() string {
	return "eth_estimateGas"
}

func (i *EthEstimateGasInput) params() interface{} {
	return []interface{}{i.Transaction.toMap()}
}

type EthEstimateGasOutput struct {
	Error  *Error `json:"error"`
	Result string `json:"result"`
}

func (i *Infura) EthEstimateGasRequest(input *EthEstimateGasInput) (req *Request, output *EthEstimateGasOutput) {
	if input == nil {
		input = &EthEstimateGasInput{}
	}

	output = &EthEstimateGasOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetBlockByHashInput struct {
	Hash   string
	IsFull bool
}

func (i *EthGetBlockByHashInput) method() string {
	return "eth_getBlockByHash"
}

func (i *EthGetBlockByHashInput) params() interface{} {
	return []interface{}{i.Hash, i.IsFull}
}

type EthGetBlockByHashOutput struct {
	Error  *Error `json:"error"`
	Result Block  `json:"result"`
}

func (i *Infura) EthGetBlockByHashRequest(input *EthGetBlockByHashInput) (req *Request, output *EthGetBlockByHashOutput) {
	if input == nil {
		input = &EthGetBlockByHashInput{}
	}

	output = &EthGetBlockByHashOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetBlockByNumberInput struct {
	BlockParameter BlockParameter
	IsFull         bool
}

func (i *EthGetBlockByNumberInput) method() string {
	return "eth_getBlockByNumber"
}

func (i *EthGetBlockByNumberInput) params() interface{} {
	return []interface{}{i.BlockParameter.Value, i.IsFull}
}

type EthGetBlockByNumberOutput struct {
	Error  *Error `json:"error"`
	Result Block  `json:"result"`
}

func (i *Infura) EthGetBlockByNumberRequest(input *EthGetBlockByNumberInput) (req *Request, output *EthGetBlockByNumberOutput) {
	if input == nil {
		input = &EthGetBlockByNumberInput{}
	}

	output = &EthGetBlockByNumberOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetTransactionByHashInput struct {
	Hash string
}

func (i *EthGetTransactionByHashInput) method() string {
	return "eth_getTransactionByHash"
}

func (i *EthGetTransactionByHashInput) params() interface{} {
	return []interface{}{i.Hash}
}

type EthGetTransactionByHashOutput struct {
	Error  *Error         `json:"error"`
	Result RawTransaction `json:"result"`
}

func (i *Infura) EthGetTransactionByHashRequest(input *EthGetTransactionByHashInput) (req *Request, output *EthGetTransactionByHashOutput) {
	if input == nil {
		input = &EthGetTransactionByHashInput{}
	}

	output = &EthGetTransactionByHashOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetTransactionByBlockHashAndIndexInput struct {
	Hash     string
	Position int
}

func (i *EthGetTransactionByBlockHashAndIndexInput) method() string {
	return "eth_getTransactionByBlockHashAndIndex"
}

func (i *EthGetTransactionByBlockHashAndIndexInput) params() interface{} {
	return []string{i.Hash, "0x" + toHexString(i.Position)}
}

type EthGetTransactionByBlockHashAndIndexOutput struct {
	Error  *Error         `json:"error"`
	Result RawTransaction `json:"result"`
}

func (i *Infura) EthGetTransactionByBlockHashAndIndexRequest(input *EthGetTransactionByBlockHashAndIndexInput) (req *Request, output *EthGetTransactionByBlockHashAndIndexOutput) {
	if input == nil {
		input = &EthGetTransactionByBlockHashAndIndexInput{}
	}

	output = &EthGetTransactionByBlockHashAndIndexOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetTransactionByBlockNumberAndIndexInput struct {
	BlockParameter BlockParameter
	Position       int
}

func (i *EthGetTransactionByBlockNumberAndIndexInput) method() string {
	return "eth_getTransactionByBlockNumberAndIndex"
}

func (i *EthGetTransactionByBlockNumberAndIndexInput) params() interface{} {
	return []string{i.BlockParameter.Value, "0x" + toHexString(i.Position)}
}

type EthGetTransactionByBlockNumberAndIndexOutput struct {
	Error  *Error         `json:"error"`
	Result RawTransaction `json:"result"`
}

func (i *Infura) EthGetTransactionByBlockNumberAndIndexRequest(input *EthGetTransactionByBlockNumberAndIndexInput) (req *Request, output *EthGetTransactionByBlockNumberAndIndexOutput) {
	if input == nil {
		input = &EthGetTransactionByBlockNumberAndIndexInput{}
	}

	output = &EthGetTransactionByBlockNumberAndIndexOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetTransactionReceiptInput struct {
	Hash string
}

func (i *EthGetTransactionReceiptInput) method() string {
	return "eth_getTransactionReceipt"
}

func (i *EthGetTransactionReceiptInput) params() interface{} {
	return []string{i.Hash}
}

type EthGetTransactionReceiptOutput struct {
	Error  *Error             `json:"error"`
	Result TransactionReceipt `json:"result"`
}

func (i *Infura) EthGetTransactionReceiptRequest(input *EthGetTransactionReceiptInput) (req *Request, output *EthGetTransactionReceiptOutput) {
	if input == nil {
		input = &EthGetTransactionReceiptInput{}
	}

	output = &EthGetTransactionReceiptOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetUncleByBlockHashAndIndexInput struct {
	Hash     string
	Position int
}

func (i *EthGetUncleByBlockHashAndIndexInput) method() string {
	return "eth_getUncleByBlockHashAndIndex"
}

func (i *EthGetUncleByBlockHashAndIndexInput) params() interface{} {
	return []string{i.Hash, "0x" + toHexString(i.Position)}
}

type EthGetUncleByBlockHashAndIndexOutput struct {
	Error  *Error `json:"error"`
	Result Block  `json:"result"`
}

func (i *Infura) EthGetUncleByBlockHashAndIndexRequest(input *EthGetUncleByBlockHashAndIndexInput) (req *Request, output *EthGetUncleByBlockHashAndIndexOutput) {
	if input == nil {
		input = &EthGetUncleByBlockHashAndIndexInput{}
	}

	output = &EthGetUncleByBlockHashAndIndexOutput{}
	req = i.newRequest(input, output)
	return
}

type EthGetUncleByBlockNumberAndIndexInput struct {
	BlockParameter BlockParameter
	Position       int
}

func (i EthGetUncleByBlockNumberAndIndexInput) method() string {
	return "eth_getUncleByBlockNumberAndIndex"
}

func (i EthGetUncleByBlockNumberAndIndexInput) params() interface{} {
	return []string{i.BlockParameter.Value, "0x" + toHexString(i.Position)}
}

type EthGetUncleByBlockNumberAndIndexOutput struct {
	Error  *Error `json:"error"`
	Result Block  `json:"result"`
}

func (i *Infura) EthGetUncleByBlockNumberAndIndexRequest(input *EthGetUncleByBlockNumberAndIndexInput) (req *Request, output *EthGetUncleByBlockNumberAndIndexOutput) {
	if input == nil {
		input = &EthGetUncleByBlockNumberAndIndexInput{}
	}

	output = &EthGetUncleByBlockNumberAndIndexOutput{}
	req = i.newRequest(input, output)
	return
}
