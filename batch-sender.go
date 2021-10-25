package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/ququzone/ethereum-scripts-go/contracts"
	"github.com/ququzone/go-common/env"
)

func main() {
	mnemonic := "desk plate waste intact rural race tired luxury hover exotic noise wonder"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	keyBytes, err := hex.DecodeString(env.GetNonEmpty("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("decode private key error: %v", err)
	}
	key, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		log.Fatalf("create esdsa private key from key bytes error: %v", err)
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(4689))
	totalAmount, _ := new(big.Int).SetString("10000000000000000000", 10)
	transactor.Value = totalAmount
	if err != nil {
		log.Fatalf("create keyed transactor error: %v", err)
	}

	client, err := ethclient.Dial("https://babel-test-106.iotex.me/")
	if err != nil {
		log.Fatalf("connect rpc error: %v", err)
	}
	contractAddr := common.HexToAddress("0x009b6cC9854030fE29F14279E0B3b60c9d657fD0")
	contract, err := contracts.NewMultiSender(contractAddr, client)
	if err != nil {
		log.Fatalf("construct contract error: %v", err)
	}

	batch, err := strconv.ParseInt(env.GetNonEmpty("BATCH"), 10, 32)
	if err != nil {
		log.Fatalf("parse batch error: %v", err)
	}

	tokenAddr := common.HexToAddress("0x6fbCdc1169B5130C59E72E51Ed68A84841C98cd1")
	tokenAmount, _ := new(big.Int).SetString("1", 10)
	coinAmount, _ := new(big.Int).SetString("200000000000000000", 10)
	for i := int(batch); i < 2000; i++ {
		addresses := make([]common.Address, 50)
		for j := 0; j < 50; j++ {
			index := i*50 + j
			path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", index))
			account, err := wallet.Derive(path, false)
			if err != nil {
				log.Fatalf("derive account error: %v", err)
			}
			addresses[j] = account.Address
		}
		tx, err := contract.Send(transactor, addresses, tokenAddr, coinAmount, tokenAmount)
		if err != nil {
			log.Fatalf("batch #d send error: %v", i, err)
		}

		time.Sleep(20 * time.Second)
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatalf("batch #d send query receipt error: %v", i, err)
		}

		if receipt.Status != 0 {
			log.Fatalf("batch #d send receipt failure: %s", i, tx.Hash().Hex())
		}
		log.Printf("batch #d transaction hash: %s\n", i, tx.Hash().Hex())
	}
}
