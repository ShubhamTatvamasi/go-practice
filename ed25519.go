package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	// create a random value from io.Reader
	var rand io.Reader

	// create public and private keys
	publicKey, privateKey, err := ed25519.GenerateKey(rand)

	// print keys
	fmt.Println(base64.StdEncoding.EncodeToString(publicKey))
	fmt.Println(base64.StdEncoding.EncodeToString(privateKey))

	// if error then print it
	if err != nil {
		fmt.Println(err)
	}

}
