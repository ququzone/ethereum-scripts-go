package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ququzone/ethereum-scripts-go/contracts"
)

func main() {
	client, _ := ethclient.Dial("https://babel-api.mainnet.iotex.io")

	multicall, _ := contracts.NewMulticall(common.HexToAddress("0x47b190bfa886788c5f974b24e192980bb89732d9"), client)

	methodId, _ := hex.DecodeString("6352211e")

	data := make(map[string][]int)
	for i := 0; i < 20; i++ {
		address := make([]common.Address, 500)
		bytecodes := make([][]byte, 500)
		s := i*500 + 1
		for j := s; j < s+500; j++ {
			address[j-s] = common.HexToAddress("0x0c5ab026d74c451376a4798342a685a0e99a5bee")
			tokenId := common.LeftPadBytes(big.NewInt(int64(j)).Bytes(), 32)
			bytecodes[j-s] = append(methodId, tokenId...)
		}

		result, _ := multicall.Batch(
			nil,
			address,
			bytecodes,
		)
		for i, v := range result.ReturnData {
			account := common.BytesToAddress(v)
			c, ok := data[account.String()]
			if !ok {
				c = make([]int, 0)
			}
			c = append(c, i+s)
			data[account.String()] = c
		}
	}

	fmt.Print("holder,count,tokenIds\n")
	for holder, v := range data {
		fmt.Printf("%s,%d,%v\n", holder, len(v), v)
	}
}
