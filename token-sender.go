package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/hex"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	keyBytes, err := hex.DecodeString(env.GetNonEmpty("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("decode private key error: %v", err)
	}
	key, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		log.Fatalf("create esdsa private key from key bytes error: %v", err)
	}
	chainId, err := strconv.ParseInt(env.GetNonEmpty("CHAIN_ID"), 10, 64)
	if err != nil {
		log.Fatalf("parse chainId error: %v", err)
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(chainId))
	if err != nil {
		log.Fatalf("create keyed transactor error: %v", err)
	}

	for i, row := range rows {
		amount, err := strconv.ParseUint(row[1], 10, 64)
		if err != nil {
			log.Fatalf("parse %d row amount error: %v", i+1, err)
		}
		value, _ := new(big.Float).Mul(big.NewFloat(math.Pow10(int(decimals))), big.NewFloat(float64(amount))).Int(nil)

		tx, err := token.Transfer(transactor, common.HexToAddress(row[0]), value)
		if err != nil {
			log.Fatalf("transfer token error: %v", err)
		}
		success := false
		for i := 0; i < 10; i++ {
			time.Sleep(20 * time.Second)
			count, err := client.PendingTransactionCount(context.Background())
			if err != nil {
				log.Fatalf("query pending transaction count for %s error: %v", tx.Hash().String(), err)
			}
			if count != 0 {
				continue
			}
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatalf("query transaction %s receipt error: %v", tx.Hash().String(), err)
			}
			if receipt.Status != 1 {
				log.Fatalf("transaction %s fail with status %d", tx.Hash().String(), receipt.Status)
			}
			success = true
		}
		if !success {
			log.Fatalf("transaction %s can't onchain", tx.Hash().String())
		}
	}
}
