package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"tz_getblockio/internal/config"
	"tz_getblockio/internal/helper"
	"tz_getblockio/internal/model"
)

const NumBlocksToInclude = 100

type MainNet interface {
	GetLastBlockNumber() (model.GetLastBlockResponse, error)
	GetBlockByNumber(addressDec int64) (model.GetBlockByNumberResponse, error)
	GetMaxChange() (model.Result, error)
}

type MainNetService struct {
	cfg    config.Config
	client http.Client
}

func NewMainNetService(cfg config.Config, client http.Client) *MainNetService {
	return &MainNetService{cfg: cfg, client: client}
}

func (m MainNetService) GetLastBlockNumber() (model.GetLastBlockResponse, error) {
	lastBlockResponse := model.GetLastBlockResponse{}
	request := model.MainNetRequest{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		ID:      "getblock.io",
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		return lastBlockResponse, err
	}
	req, err := http.NewRequest("POST", "https://eth.getblock.io/mainnet/", bytes.NewReader(requestData))

	if err != nil {
		return lastBlockResponse, err
	}
	req.Header.Add("x-api-key", m.cfg.Client.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := m.client.Do(req)
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&lastBlockResponse)
	if err != nil {
		return lastBlockResponse, err
	}
	return lastBlockResponse, nil
}

func (m MainNetService) GetBlockByNumber(addressDec int64) (model.GetBlockByNumberResponse, error) {
	result := model.GetBlockByNumberResponse{}
	request := model.MainNetRequest{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		ID:      "getblock.io",
	}
	addressHex := "0x" + strconv.FormatInt(addressDec, 16)
	request.Params = append(request.Params, addressHex, true)
	requestData, err := json.Marshal(request)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("POST", "https://eth.getblock.io/mainnet/", bytes.NewReader(requestData))
	if err != nil {
		return result, err
	}
	req.Header.Add("x-api-key", m.cfg.Client.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := m.client.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m MainNetService) GetMaxChange() (model.Result, error) {
	maxValueResult := model.Result{
		In:     "",
		Out:    "",
		Amount: big.NewInt(0),
	}
	lastBlockNumber, err := m.GetLastBlockNumber()
	lastBlockNumberDec, err := strconv.ParseInt(helper.HexaNumberToInteger(lastBlockNumber.Result), 16, 32)
	if err != nil {
		return maxValueResult, err
	}

	firstBlockNumberDec := lastBlockNumberDec - NumBlocksToInclude
	var wg sync.WaitGroup
	ch := make(chan model.Result)
	for i := firstBlockNumberDec; i < lastBlockNumberDec; i++ {
		wg.Add(1)
		go func(i int64) {
			err := m.getMaxValueResultFromBlock(i, &wg, ch)
			if err != nil {
				fmt.Println(err)
			}
		}(i)
	}

	var results []model.Result
	go func() {
		for v := range ch {
			results = append(results, v)
		}
	}()

	wg.Wait()
	close(ch)

	for _, v := range results {
		if v.Amount.Cmp(maxValueResult.Amount) == 1 {
			maxValueResult = v
		}
	}

	return maxValueResult, nil
}

func (m MainNetService) getMaxValueResultFromBlock(address int64, wg *sync.WaitGroup, ch chan<- model.Result) error {
	defer wg.Done()
	result := model.Result{
		In:     "",
		Out:    "",
		Amount: big.NewInt(0),
	}
	blockData, err := m.GetBlockByNumber(address)
	if err != nil {
		return err
	}

	for _, transaction := range blockData.Result.Transactions {
		valueDec := new(big.Int)
		valueDec.SetString(helper.HexaNumberToInteger(transaction.Value), 16)
		if err != nil {
			return err
		}
		if valueDec.Cmp(result.Amount) == 1 {
			result.Amount = valueDec
			result.In = transaction.From
			result.Out = transaction.To
		}
	}
	ch <- result
	return nil
}
