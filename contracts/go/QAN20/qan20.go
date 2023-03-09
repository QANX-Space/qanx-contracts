/*
## Example

```go

package main

import (
	"os"

	qan20 "qanx.space/qanx-contracts/go/QAN20"
	callintrpr "qanx.space/qanx-contracts/go/utils/CallInterpreter"
)

func main() {
	token := qan20.NewQAN20("Example", "XMPL", 8)

	callintrpr.Interpret(token, os.Args[1:])
}

```*/
package qan20

import (
	"fmt"
	"os"
	"strconv"

	context "qanx.space/qanx-contracts/go/utils/Context"
	db "qanx.space/qanx-contracts/go/utils/Database"
)

type QAN20Token interface {
	Name() string
	Symbol() string
	Decimals() uint64
	TotalSupply() uint64
	BalanceOf(owner string) uint64
	Transfer(to string, amount uint64) bool
	Allowance(owner string, spender string) uint64
	Approve(to string, amount uint64) bool
	TransferFrom(from string, to string, amount uint64) bool
	Mint(to string, tokenId uint64)
}

// QAN20 smart contract standard
type QAN20 struct {
	// name of the contract
	name string
	// symbol of the contract
	symbol string
	// decimals of the contract
	decimals uint64
}

// Creates the QAN20 smart contract
func NewQAN20(name string, symbol string, decimals uint64) *QAN20 {
	return &QAN20{name: name, symbol: symbol, decimals: decimals}
}

// Retrieve the name
func (token *QAN20) Name() string {
	return token.name
}

// Retrieve the symbol/ticker
func (token *QAN20) Symbol() string {
	return token.symbol
}

// Retrieve the decimal amount
func (token *QAN20) Decimals() uint64 {
	return token.decimals
}

// Retrieve the total supply
func (token *QAN20) TotalSupply() uint64 {
	n, _ := strconv.ParseUint(db.Read("TOTAL_SUPPLY"), 10, 64)
	return n
}

// Retrieve the balance of owner
func (token *QAN20) BalanceOf(owner string) uint64 {
	n, _ := strconv.ParseUint(db.Read(fmt.Sprintf("BALANCE_OF_%v", owner)), 10, 64)
	return n
}

// Transfers tokens from sender to "to"
func (token *QAN20) Transfer(to string, amount uint64) bool {
	sender := context.Sender()
	return token.transfer(sender, to, amount)
}

// Retrieve the remaining amount of tokens allowed to spend by spender
func (token *QAN20) Allowance(owner string, spender string) uint64 {
	n, _ := strconv.ParseUint(db.Read(fmt.Sprintf("TOKEN_ALLOWANCE_%v_%v", owner, spender)), 10, 64)
	return n
}

// Sets amount as the allowance of spender over the caller's tokens
func (token *QAN20) Approve(spender string, amount uint64) bool {
	sender := context.Sender()
	db.Write(fmt.Sprintf("TOKEN_ALLOWANCE_%v_%v", sender, spender), strconv.FormatUint(amount, 10))
	return true
}

// Transfers tokens from "from" to "to"
func (token *QAN20) TransferFrom(from string, to string, amount uint64) bool {
	sender := context.Sender()
	allowance := token.Allowance(from, sender)

	if allowance < amount {
		os.Stderr.WriteString("QAN20: Insufficient allowance\n")
		os.Exit(1)
	}

	db.Write(fmt.Sprintf("TOKEN_ALLOWANCE_%v_%v", from, sender), strconv.FormatUint(allowance-amount, 10))

	return token.transfer(from, to, amount)
}

// Mints the amount of tokens and transfers it to "to"
func (token *QAN20) Mint(to string, amount uint64) {
	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), strconv.FormatUint(token.BalanceOf(to)+amount, 10))
	db.Write("TOTAL_SUPPLY", strconv.FormatUint(token.TotalSupply()+amount, 10))
}

func (token *QAN20) transfer(from string, to string, amount uint64) bool {
	fromBalance := token.BalanceOf(from)

	if fromBalance < amount {
		os.Stderr.WriteString("QAN20: Transfer amount exceeds balance\n")
		os.Exit(1)
	}

	db.Write(fmt.Sprintf("BALANCE_OF_%v", from), strconv.FormatUint(fromBalance-amount, 10))
	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), strconv.FormatUint(token.BalanceOf(to)+amount, 10))

	return true
}
