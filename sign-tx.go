package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ququzone/go-common/env"
)

func main() {
	tx := types.NewTransaction(
		0,
		common.HexToAddress("0x173553c179bbf5af39D8Db41F0B60e4Fc631066a"),
		big.NewInt(100),
		10000,
		big.NewInt(1000000000000),
		[]byte{},
	)

	keyBytes, err := hex.DecodeString(env.GetNonEmpty("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("decode private key error: %v", err)
	}
	privateKey, err := crypto.ToECDSA(keyBytes)
	signer := types.NewEIP155Signer(big.NewInt(4690))

	rlpTx, _ := rlp.EncodeToBytes([]interface{}{
		tx.Nonce(),
		tx.GasPrice(),
		tx.Gas(),
		tx.To(),
		tx.Value(),
		tx.Data(),
		big.NewInt(4690), uint(0), uint(0),
	})
	fmt.Printf("rlp: %s\n", hex.EncodeToString(rlpTx))
	hash := signer.Hash(tx)
	fmt.Printf("sighash: %s\n", hash.String())

	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hash: %s\n", signedTx.Hash())
	fmt.Printf("raw: %s\n", hexutil.Encode(raw))
}
