package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/Finschia/finschia-sdk/client/input"
	"github.com/Finschia/finschia-sdk/codec"
	"github.com/Finschia/finschia-sdk/codec/legacy"
	"github.com/Finschia/finschia-sdk/crypto"
	cryptocodec "github.com/Finschia/finschia-sdk/crypto/codec"
	"github.com/Finschia/finschia-sdk/crypto/keys/secp256k1"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	privKeyHex, err := input.GetPassword("Enter 64 character raw hex private secp256k1 key:", buf)
	passphrase, err := input.GetPassword("Pick a password, at least 8 chars:", buf)
	if err != nil {
		panic(err)
	}

	privKeyRaw, err := hex.DecodeString(privKeyHex)
	if err != nil {
		panic(err)
	}
	cdc := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(cdc)
	privKeyBytes := cdc.MustMarshal(secp256k1.PrivKey{Key: privKeyRaw})
	privKey, err := legacy.PrivKeyFromBytes(privKeyBytes)
	if err != nil {
		panic(err)
	}
	text := crypto.EncryptArmorPrivKey(privKey, passphrase, "")
	fmt.Println(text)
}
