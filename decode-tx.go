package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	rawTxData, err := hex.DecodeString("f8668085e8d4a5100082271094173553c179bbf5af39d8db41f0b60e4fc631066a64808224c8a0bdf6681fb6ff949fb67c07950108f0394b9ecbc991626d5e9c0386eaf5c53f01a053588e443deb34eac3b677bdded0204863df7c0d4aba0be2c1c6d4ccda3cb880")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	var tx types.Transaction
	err = rlp.DecodeBytes(rawTxData, &tx)

	fmt.Println(tx.Nonce())
	fmt.Println(tx.ChainId())
	fmt.Println(tx.To())
	fmt.Println(tx.Value())
	fmt.Println(tx.Gas())
	fmt.Println(tx.GasPrice())
	fmt.Println(tx.Data())
}
