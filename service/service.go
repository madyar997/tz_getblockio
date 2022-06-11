package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tz_getblockio/model"
)

type MainNet interface {
	GetLastBlockNumber() (model.GetLastBlockResponse, error)
	GetBlockByNumber() error
}

type MainNetService struct {
	client http.Client
}

func NewMainNetService(client http.Client) *MainNetService {
	return &MainNetService{client: client}
}

func (m MainNetService) GetLastBlockNumber() (model.GetLastBlockResponse, error) {
	result := model.GetLastBlockResponse{}
	request := model.MainNetRequest{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  nil,
		ID:      "getblock.io",
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("POST", "https://eth.getblock.io/mainnet/", bytes.NewReader(requestData))

	if err != nil {
		return result, err
	}
	req.Header.Add("x-api-key", "74f75a54-1efa-4105-a88b-54cde98aa270")
	req.Header.Add("Content-Type", "application/json")

	res, err := m.client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m MainNetService) GetBlockByNumber() error {
	return nil
}

func (m MainNetService) GetMaxChange() (string, error) {
	lastBlockNumber, err := m.GetLastBlockNumber()
	if err != nil {
		return "", err
	}
	fmt.Println(lastBlockNumber)
	return "", nil
}
