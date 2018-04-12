package infura_test

import (
	"flag"
	"testing"

	infura "github.com/daisuke310vvv/infura-go"
)

var apiKey = flag.String("apiKey", "", "api key for call infura.io APIs")

func TestApi(t *testing.T) {
	if apiKey == nil || *apiKey == "" {
		t.Fatalf("No apiKey found.")
	}

	cli, err := infura.New(infura.NewConfig(*apiKey, infura.Ropsten))
	if err != nil {
		t.Fatalf(err.Error())
	}

	req, resp := cli.EthBlockNumberRequest(nil)
	err = req.Call()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("block number: %v", resp.Result)
}
