package main

import (
	"bufio"
	"context"
	"encoding/hex"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ququzone/ethereum-scripts-go/contracts"
	"github.com/ququzone/go-common/env"
)

func main() {
	client, err := ethclient.Dial("https://babel-api.mainnet.iotex.io")
	if err != nil {
		log.Fatalf("connect rpc error: %v", err)
	}
	contractAddr := common.HexToAddress("0x34C3Ad42767b758665106D8fbeA346CbeC7a47aa")
	contract, err := contracts.NewMultiSender(contractAddr, client)
	if err != nil {
		log.Fatalf("construct contract error: %v", err)
	}

	file, err := os.Open("./addresses.txt")
	if err != nil {
		log.Fatalf("failed opening buckets file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	addresses := make([]common.Address, 0)
	for scanner.Scan() {
		row := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		addresses = append(addresses, common.HexToAddress(row[0]))
	}

	times := int(len(addresses) / 50)
	if len(addresses)%50 != 0 {
		times++
	}

	zeroAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")
	iotx, _ := new(big.Int).SetString("20000000000000000000", 10)

	keyBytes, err := hex.DecodeString(env.GetNonEmpty("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("decode private key error: %v", err)
	}
	key, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		log.Fatalf("create esdsa private key from key bytes error: %v", err)
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(4689))
	if err != nil {
		log.Fatalf("create keyed transactor error: %v", err)
	}

	for i := 0; i < times; i++ {
		size := 50
		if i == times-1 && len(addresses)%50 != 0 {
			size = int(len(addresses) % 50)
		}
		addrs := make([]common.Address, size)
		for j := 0; j < size; j++ {
			addrs[j] = addresses[i*50+j]
		}

		totalAmount := new(big.Int).Mul(iotx, big.NewInt(int64(size)))
		transactor.Value = totalAmount

		tx, err := contract.Send(transactor, addrs, zeroAddr, iotx, big.NewInt(0))
		if err != nil {
			log.Fatalf("batch %d send error: %v", i, err)
		}

		time.Sleep(20 * time.Second)
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatalf("batch %d send query receipt error: %v", i, err)
		}

		if receipt.Status != 1 {
			log.Fatalf("batch %d send receipt failure: %s", i, tx.Hash().Hex())
		}
		log.Printf("batch %d transaction hash: %s\n", i, tx.Hash().Hex())
	}
}
