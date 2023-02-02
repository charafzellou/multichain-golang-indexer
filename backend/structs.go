package main

type ContractCode struct {
	Code string `json:"code"`
}

type ContractData struct {
	Operations []Operation
}

type Operation struct {
	Timestamp int32   `json:"timestamp"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Hash      string  `json:"hash"`
	Value     float64 `json:"value"`
	Input     string  `json:"input"`
	Success   bool    `json:"success"`
}

type Operations struct {
	OperationList []Operation `json:"operations"`
}

type TokenOperation struct {
	Timestamp       int32       `json:"timestamp"`
	TransactionHash string      `json:"transactionHash"`
	TokenInfo       interface{} `json:"tokenInfo"`
	Type            string      `json:"type"`
	Address         string      `json:"address"`
	From            string      `json:"from"`
	To              string      `json:"to"`
	Value           float64     `json:"value"`
}

type TokenOperations struct {
	OperationList []TokenOperation `json:"operations"`
}

type Address struct {
	Address      string       `json:"address"`
	ETH          ETH          `json:"ETH"`
	ContractInfo ContractInfo `json:"contractInfo,omitempty"`
	TokenInfo    TokenInfo    `json:"tokenInfo,omitempty"`
	Tokens       []Token      `json:"tokens,omitempty"`
	CountTxs     int32        `json:"countTxs"`
}

type ETH struct {
	Balance    float64 `json:"balance"`
	RawBalance string  `json:"rawBalance"`
	TotalIn    int32   `json:"totalIn"`
	TotalOut   int32   `json:"totalOut"`
}

type ContractInfo struct {
	CreatorAddress  string `json:"creatorAddress"`
	TransactionHash string `json:"transactionHash"`
	Timestamp       int32  `json:"timestamp"`
}

type TokenInfo struct {
	// TokenInfo data
}

type Token struct {
	TokenInfo  TokenInfo `json:"tokenInfo"`
	RawBalance string    `json:"rawBalance"`
	TotalIn    int32     `json:"totalIn"`
	TotalOut   int32     `json:"totalOut"`
}

type OperationByTimestamp []Operation

func (a OperationByTimestamp) Len() int           { return len(a) }
func (a OperationByTimestamp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OperationByTimestamp) Less(i, j int) bool { return a[i].Timestamp > a[j].Timestamp }

type TokenOperationByTimestamp []TokenOperation

func (a TokenOperationByTimestamp) Len() int           { return len(a) }
func (a TokenOperationByTimestamp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TokenOperationByTimestamp) Less(i, j int) bool { return a[i].Timestamp > a[j].Timestamp }
