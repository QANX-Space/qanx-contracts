package block

import (
	"math/big"
	"os"

	common "qanx.space/qanx-contracts/go/utils/Common"
)

// Retrieves the prevHash of the block
func PrevHash() *big.Int {
	n, _ := common.ParseBig256(os.Getenv("BLOCK_PREVHASH"))

	return n
}

// Retrieves the number of the block
func Number() *big.Int {
	n, _ := common.ParseBig256(os.Getenv("BLOCK_NUMBER"))

	return n
}

// Retrieves the chain id of the block
func ChainId() *big.Int {
	n, _ := common.ParseBig256(os.Getenv("BLOCK_CHAINID"))

	return n
}

// Retrieves the coinbase of the block
func Coinbase() *big.Int {
	n, _ := common.ParseBig256(os.Getenv("BLOCK_COINBASE"))

	return n
}

// Retrieves the difficulty of the block
func Difficulty() *big.Int {
	n, _ := common.ParseBig256(os.Getenv("BLOCK_DIFFICULTY"))

	return n
}

// Retrieves the gasLimit of the block
func GasLimit() *big.Int {
	n, _ := common.ParseBig256(os.Getenv("BLOCK_GASLIMIT"))

	return n
}

// Retrieves the timestamp of the block
func Timestamp() *big.Int {
	n, _ := common.ParseBig256(os.Getenv("BLOCK_TIMESTAMP"))

	return n
}
