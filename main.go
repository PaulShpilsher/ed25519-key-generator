package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	fileName := flag.String("out", "key.ed25519", "output key file name")
	flag.Parse()

	privateKeyFileName := fmt.Sprintf("%s.prv", *fileName)
	publicKeyFileName := fmt.Sprintf("%s.pub", *fileName)

	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		fmt.Println("failed to generate key pair:", err)
		return
	}
	writeFile(privateKeyFileName, privateKey)

	writeFile(publicKeyFileName, publicKey)

	fmt.Println("Files:", privateKeyFileName, publicKeyFileName)
	fmt.Println("Public Key:", hex.EncodeToString(publicKey))
	fmt.Println("Private Key:", hex.EncodeToString(privateKey))
}

func writeFile(fileName string, data []byte) {
	if err := os.WriteFile(fileName, data, 0600); err != nil {
		log.Fatalf("writing \"%s\" file failed. err: %v\n", fileName, err)
	}
}
