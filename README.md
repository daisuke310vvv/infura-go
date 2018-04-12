# infura-go
[INFURA](https://infura.io) api client written in golang

## Installation  
To install `infura-go`, simply execute the following command in a terminal from your `$GOPATH`:

```
go get github.com/daisuke310vvv/infura-go
```

## How to use  

```
config := infura.NewConfig("<INFURA API KEY>", infura.Ropsten)
infuraClient := infura.New(config)

input := &infura.EthGetBalanceInput{
    Address:        "0x5c66b0d82df26e8FE165Be6628F5f5e1f1bccD5C",
    BlockParameter: infura.NewBlockParameter("latest"),
}
req, resp := infuraClient.EthGetBalanceRequest(input)
_ = req.Call()

fmt.Println(res.Result)
// => 0x7759d50972e9800
```

## Test  
**NOTE**: Be careful many api calls will be restricted.  

```
$ go test *_test.go -apiKey="<INFURA API KEY>"
```

## TODO  
- [ ] Error handling
- [ ] Decode all result parameters
- [ ] Test

## Supported JSON-RPC methods  

- `web3_clientVerwsion`
- `web3_sha3`
- `net_version`
- `net_listening`
- `net_peerCount`
- `eth_protocolVersion`
- `eth_syncing`
- `eth_mining`
- `eth_hashrate`
- `eth_gasPrice`
- `eth_accounts`
- `eth_blockNumber`
- `eth_getBalance`
- `eth_getTransactionCount`
- `eth_getBlockTransactionCountByHash`
- `eth_getBlockTransactionCountByNumber`
- `eth_getUncleCountByBlockHash`
- `eth_getUncleCountByBlockNumber`
- `eth_getCode`
- `eth_sendRawTransaction`
- `eth_call`
- `eth_estimateGas`
- `eth_getBlockByHash`
- `eth_getBlockByNumber`
- `eth_getTransactionByHash`
- `eth_getTransactionByBlockHashAndIndex`
- `eth_getTransactionByBlockNumberAndIndex`
- `eth_getTransactionReceipt`
- `eth_getUncleByBlockHashAndIndex`
- `eth_getUncleByBlockNumberAndIndex`

## Donation  
Donations help me keep this development❤️  
Send me eth to `0x5c66b0d82df26e8FE165Be6628F5f5e1f1bccD5C`


## License  
`infura-go` is available under the MIT license. See the LICENSE file for more info.
