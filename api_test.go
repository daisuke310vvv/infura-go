package infura_test

import (
	"flag"
	"testing"

	infura "github.com/daisuke310vvv/infura-go"
)

var projectID = flag.String("projectID", "", "project id for call infura.io APIs")

func TestApi(t *testing.T) {
	if projectID == nil || *projectID == "" {
		t.Fatalf("No projectID found.")
	}

	cli, err := infura.New(infura.NewConfig(*projectID, infura.Ropsten))
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
