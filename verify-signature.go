package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	hash, _ := hex.DecodeString("78fbc2d79c9011e62c86a5ffed115fa5280420ffbcc05a550c57e7482617f807")
	// d20153d2ba1e66e7e0b42ef58ab545a7d83c6ae1e995cdf888ad1c0f8151d20708c2307b83d9cbcc715383240d55bb24189f1ec8752aaed5c175620433b0cc5f01
	// 81171b1b59f6714a620d2591e3084c106b6d41a1c11abb0954df85142688196864e4abbe434e57b35d69db80f314d23b649e63e4669edf3f0fa57a0de9ec2f4701
	signature, _ := hex.DecodeString("81171b1b59f6714a620d2591e3084c106b6d41a1c11abb0954df85142688196864e4abbe434e57b35d69db80f314d23b649e63e4669edf3f0fa57a0de9ec2f4701")

	sigPublicKey, err := crypto.Ecrecover(hash, signature)
	if err != nil {
		log.Fatal(err)
	}

	pubkey, _ := crypto.UnmarshalPubkey(sigPublicKey)
	address := crypto.PubkeyToAddress(*pubkey)

	fmt.Println(hex.EncodeToString(sigPublicKey))
	fmt.Println(address)
}
