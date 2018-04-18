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

	req, resp := cli.EthGetBalanceRequest(&infura.EthGetBalanceInput{Address: "0x5c66b0d82df26e8FE165Be6628F5f5e1f1bccD5C", BlockParameter: infura.NewBlockParameter("latest")})
	err = req.Call()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if resp.Error != nil {
		t.Fatalf(resp.Error.Error())
	}

	t.Logf("balance: %v", resp.Result)
}
