package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://api.nightly-cluster-2.iotex.one:15014")
	if err != nil {
		log.Fatalf("connect rpc error: %v", err)
	}
}
