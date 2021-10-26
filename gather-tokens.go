package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/ququzone/ethereum-scripts-go/contracts"
)

func main() {
	mnemonic := "desk plate waste intact rural race tired luxury hover exotic noise wonder"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("https://babel-test-106.iotex.me/")
	if err != nil {
		log.Fatalf("connect rpc error: %v", err)
	}

	token, err := contracts.NewERC20(common.HexToAddress("0x6fbCdc1169B5130C59E72E51Ed68A84841C98cd1"), client)
	if err != nil {
		log.Fatalf("create token contract error: %v", err)
	}
	for i := 0; i < 100000; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
		account, err := wallet.Derive(path, false)
		if err != nil {
			log.Fatalf("derive account %d error: %v", i, err)
		}
		key, err := wallet.PrivateKey(account)
		if err != nil {
			log.Fatalf("get account %d private key error: %v", i, err)
		}
		transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(4689))
		tx, err := token.Transfer(transactor, common.HexToAddress("0x8896780a7912829781f70344ab93e589dddb2930"), big.NewInt(1))
		if err != nil {
			log.Fatalf("send account %d transfer tx error: %v", i, err)
		}
		log.Printf("send account %d transfer transaction hash %s", i, tx.Hash().Hex())
	}
}
