package model

import "math/big"

type MainNetRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      string        `json:"id"`
}

type Param struct {
	Address string
	Bool    bool
}

type GetLastBlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result"`
}

type GetBlockByNumberResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		BaseFeePerGas   string `json:"baseFeePerGas"`
		Difficulty      string `json:"difficulty"`
		ExtraData       string `json:"extraData"`
		GasLimit        string `json:"gasLimit"`
		GasUsed         string `json:"gasUsed"`
		Hash            string `json:"hash"`
		LogsBloom       string `json:"logsBloom"`
		Miner           string `json:"miner"`
		MixHash         string `json:"mixHash"`
		Nonce           string `json:"nonce"`
		Number          string `json:"number"`
		ParentHash      string `json:"parentHash"`
		ReceiptsRoot    string `json:"receiptsRoot"`
		Sha3Uncles      string `json:"sha3Uncles"`
		Size            string `json:"size"`
		StateRoot       string `json:"stateRoot"`
		Timestamp       string `json:"timestamp"`
		TotalDifficulty string `json:"totalDifficulty"`
		Transactions    []struct {
			BlockHash            string        `json:"blockHash"`
			BlockNumber          string        `json:"blockNumber"`
			From                 string        `json:"from"`
			Gas                  string        `json:"gas"`
			GasPrice             string        `json:"gasPrice"`
			Hash                 string        `json:"hash"`
			Input                string        `json:"input"`
			Nonce                string        `json:"nonce"`
			To                   string        `json:"to"`
			TransactionIndex     string        `json:"transactionIndex"`
			Value                string        `json:"value"`
			Type                 string        `json:"type"`
			V                    string        `json:"v"`
			R                    string        `json:"r"`
			S                    string        `json:"s"`
			MaxFeePerGas         string        `json:"maxFeePerGas,omitempty"`
			MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas,omitempty"`
			AccessList           []interface{} `json:"accessList,omitempty"`
			ChainID              string        `json:"chainId,omitempty"`
		} `json:"transactions"`
		TransactionsRoot string   `json:"transactionsRoot"`
		Uncles           []string `json:"uncles"`
	} `json:"result"`
}

type Result struct {
	In     string   `json:"in"`
	Out    string   `json:"out"`
	Amount *big.Int `json:"amount"`
}
