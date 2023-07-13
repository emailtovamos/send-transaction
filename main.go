package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	// Connect to the local Ethereum network (replace "http://localhost:8545" with your local network endpoint)
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	// Load the private key of the account that will send the transaction
	privateKey, err := crypto.HexToECDSA("9b28f36fbd67381120752d6172ecdcf10e06ab2d9a1367aac00cdcd6ac7855d3")
	if err != nil {
		log.Fatal(err)
	}

	// Get the public key and Ethereum address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the nonce for the sender's address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare the transaction data
	// Replace "0x123456789ABCDEF" with the recipient address
	// Replace "1000000000000000000" with the amount to send (in Wei)
	// Replace "200000" with the gas limit
	// Replace "1000000000" with the gas price (in Wei)
	toAddress := common.HexToAddress("0x63FaC9201494f0bd17B9892B9fae4d52fe3BD377")
	amount := big.NewInt(1000000000000000000)
	gasLimit := uint64(200000)
	gasPrice := big.NewInt(1000000000)
	data := []byte{} // Optional: Add any data you want to include in the transaction

	// Create the transaction
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, data)

	// // Sign the transaction with the sender's private key
	// chainID, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	fmt.Println("error 1")
	// 	log.Fatal(err)
	// }
	
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1337)), privateKey)
	if err != nil {
		fmt.Println("error 2")
		log.Fatal(err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("error 3")
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
}
