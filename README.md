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

req, resp := infuraClient.EthBlockNumberRequest(nil)
_ = req.Call()

fmt.Println(res.Result)
// => 0x2e2b9a
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

## License  
`infura-go` is available under the MIT license. See the LICENSE file for more info.
