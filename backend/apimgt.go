package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func getContract(contractAddr string) []string {
	url := fmt.Sprintf("%s/getAddressInfo/%s?apiKey=%s", apiMainnet, contractAddr, apiKey)
	fmt.Printf("%s/getAddressInfo/%s?apiKey=%s\n", apiMainnet, contractAddr, apiKey)

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var adresseInfo Address
	err = json.Unmarshal(body, &adresseInfo)
	if err != nil {
		panic(err)
	}

	printContract(adresseInfo)
	return genContractArray(adresseInfo)
}

func printContract(adresseInfo Address) {
	fmt.Printf("Address %s:\n", adresseInfo.Address)
	fmt.Printf("- ETH Balance: %f\n", adresseInfo.ETH.Balance)
	fmt.Printf("- Creator: %s\n", adresseInfo.ContractInfo.CreatorAddress)
	fmt.Printf("- Creation Timestamp: %d\n", adresseInfo.ContractInfo.Timestamp)
	fmt.Printf("- Transactions: %d\n", adresseInfo.CountTxs)
}

func genContractArray(adresseInfo Address) []string {
	// address TEXT, eth_balance FLOAT, creator TEXT, creation_timestamp BIGINT, transactions BIGINT
	var adresseSqlArray []string
	adresseSqlArray = append(adresseSqlArray, adresseInfo.Address)
	adresseSqlArray = append(adresseSqlArray, strconv.FormatFloat(adresseInfo.ETH.Balance, 'f', -1, 64))
	adresseSqlArray = append(adresseSqlArray, adresseInfo.ContractInfo.CreatorAddress)
	adresseSqlArray = append(adresseSqlArray, strconv.Itoa(int(adresseInfo.ContractInfo.Timestamp)))
	adresseSqlArray = append(adresseSqlArray, strconv.Itoa(int(adresseInfo.CountTxs)))
	for index, value := range adresseSqlArray {
		fmt.Println("Index:", index, "Value:", value)
	}
	return adresseSqlArray
}

func getOperations(contractAddr string) [][]string {
	url := fmt.Sprintf("%s/getAddressTransactions/%s?apiKey=%s&limit=%d&showZeroValues=%t", apiMainnet, contractAddr, apiKey, apiLimit, apiShowZeroValues)

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var operations []Operation
	err = json.Unmarshal(body, &operations)
	if err != nil {
		panic(err)
	}

	sort.Sort(OperationByTimestamp(operations))
	printOperations(operations)
	return genOperationsArray(operations)
}

func genOperationsArray(operations []Operation) [][]string {
	// hash TEXT, timestamp BIGINT, from TEXT, to TEXT, value FLOAT, input TEXT
	var operationsSqlArray [][]string
	for _, op := range operations {
		var opSqlArray []string
		opSqlArray = append(opSqlArray, op.Hash)
		opSqlArray = append(opSqlArray, strconv.Itoa(int(op.Timestamp)))
		opSqlArray = append(opSqlArray, op.From)
		opSqlArray = append(opSqlArray, op.To)
		opSqlArray = append(opSqlArray, strconv.FormatFloat(op.Value, 'f', -1, 64))
		opSqlArray = append(opSqlArray, op.Input)
		opSqlArray = append(opSqlArray, strconv.FormatBool(op.Success))
		operationsSqlArray = append(operationsSqlArray, opSqlArray)
	}
	return operationsSqlArray
}

func printOperations(operations []Operation) {
	for i, operation := range operations {
		fmt.Printf("Operation %d:\n", i)
		fmt.Printf("- Timestamp: %s\n", (time.Unix(int64(operation.Timestamp), 0)).Format(time.RFC3339))
		fmt.Printf("- From: %s\n", operation.From)
		fmt.Printf("- To: %s\n", operation.To)
		fmt.Printf("- Hash: %s\n", operation.Hash)
		fmt.Printf("- Value: %f ETH\n", operation.Value)
		fmt.Printf("- Input: %s\n", operation.Input)
		fmt.Printf("- Success: %t\n", operation.Success)
	}
}
