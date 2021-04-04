package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ququzone/ethereum-scripts-go/contracts"
	"github.com/ququzone/go-common/env"
)

func main() {
	contentBytes, err := ioutil.ReadFile("./token-sender-list.csv")
	if err != nil {
		log.Fatalf("read csv file error: %v", err)
	}
	r := csv.NewReader(bytes.NewReader(contentBytes))
	rows, _ := r.ReadAll()

	client, err := ethclient.Dial(env.GetNonEmpty("ETH_ENDPOINT"))
	if err != nil {
		log.Fatalf("connect eth endpoint error: %v", err)
	}
	token, err := contracts.NewERC20(common.HexToAddress(env.GetNonEmpty("TOKEN_ADDRESSS")), client)
	if err != nil {
		log.Fatalf("create %s contract error: %v", env.GetNonEmpty("TOKEN_ADDRESSS"), err)
	}
	decimals, err := token.Decimals(nil)
	if err != nil {
		log.Fatalf("fetch %s token decimals error: %v", env.GetNonEmpty("TOKEN_ADDRESSS"), err)
	}

	for i, row := range rows {
		amount, err := strconv.ParseUint(row[1], 10, 64)
		if err != nil {
			log.Fatalf("parse %d row amount error: %v", i+1, err)
		}
		value, _ := new(big.Float).Mul(big.NewFloat(math.Pow10(int(decimals))), big.NewFloat(float64(amount))).Int(nil)

		tx, err := token.Transfer(nil, common.HexToAddress(row[0]), value)

		fmt.Printf("%s:%s\n", row[0], tokenAmount.String())
	}
}
