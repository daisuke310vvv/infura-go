package infura

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Request *http.Request
	Error   error
	Data    interface{}
}

func (r *Request) Call() error {
	if r.Error != nil {
		return r.Error
	}

	resp, err := http.DefaultClient.Do(r.Request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, r.Data)
	if err != nil {
		return err
	}

	return nil
}

func (i *Infura) newRequest(input input, output interface{}) (req *Request) {
	args := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  input.method(),
		"params":  input.params(),
	}

	body, err := json.Marshal(args)
	if err != nil {
		body = []byte{}
	}

	httpReq, err := http.NewRequest("POST", i.network.URL()+i.apiKey, bytes.NewReader(body))

	req = &Request{
		Request: httpReq,
		Error:   err,
		Data:    output,
	}

	return
}
