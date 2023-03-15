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
	"math/big"
	"os"
	"strings"

	common "qanx.space/qanx-contracts/go/utils/Common"
	db "qanx.space/qanx-contracts/go/utils/Database"
	message "qanx.space/qanx-contracts/go/utils/Message"
)

type QAN20Token interface {
	Name() string
	Symbol() string
	Decimals() *big.Int
	TotalSupply() *big.Int
	BalanceOf(owner string) *big.Int
	Transfer(to string, amount *big.Int) bool
	Allowance(owner string, spender string) *big.Int
	Approve(to string, amount *big.Int) bool
	TransferFrom(from string, to string, amount *big.Int) bool
	Mint(to string, amount *big.Int)
}

// QAN20 smart contract standard
type QAN20 struct {
	// name of the contract
	name string
	// symbol of the contract
	symbol string
	// decimals of the contract
	decimals *big.Int
}

// Creates the QAN20 smart contract
func NewQAN20(name string, symbol string, decimals *big.Int) *QAN20 {
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
func (token *QAN20) Decimals() *big.Int {
	return token.decimals
}

// Retrieve the total supply
func (token *QAN20) TotalSupply() *big.Int {
	n, _ := common.ParseBig256(db.Read("TOTAL_SUPPLY"))
	return n
}

// Retrieve the balance of owner
func (token *QAN20) BalanceOf(owner string) *big.Int {
	owner = strings.ToLower(owner)
	n, _ := common.ParseBig256(db.Read(fmt.Sprintf("BALANCE_OF_%v", owner)))

	return n
}

// Transfers tokens from sender to "to"
func (token *QAN20) Transfer(to string, amount *big.Int) bool {
	to = strings.ToLower(to)
	sender := message.Sender()

	return token.transfer(sender, to, amount)
}

// Retrieve the remaining amount of tokens allowed to spend by spender
func (token *QAN20) Allowance(owner string, spender string) *big.Int {
	owner = strings.ToLower(owner)
	spender = strings.ToLower(spender)

	n, _ := common.ParseBig256(db.Read(fmt.Sprintf("TOKEN_ALLOWANCE_%v_%v", owner, spender)))
	return n
}

// Sets amount as the allowance of spender over the caller's tokens
func (token *QAN20) Approve(spender string, amount *big.Int) bool {
	spender = strings.ToLower(spender)
	sender := message.Sender()

	db.Write(fmt.Sprintf("TOKEN_ALLOWANCE_%v_%v", sender, spender), amount.String())
	return true
}

// Transfers tokens from "from" to "to"
func (token *QAN20) TransferFrom(from string, to string, amount *big.Int) bool {
	from = strings.ToLower(from)
	to = strings.ToLower(to)

	sender := message.Sender()
	allowance := token.Allowance(from, sender)

	if allowance.Cmp(amount) == -1 /* Allowance is lower than amount */ {
		os.Stderr.WriteString("QAN20: Insufficient allowance\n")
		os.Exit(1)
	}

	db.Write(fmt.Sprintf("TOKEN_ALLOWANCE_%v_%v", from, sender), allowance.Sub(allowance, amount).String())

	return token.transfer(from, to, amount)
}

// Mints the amount of tokens and transfers it to "to"
func (token *QAN20) Mint(to string, amount *big.Int) {
	to = strings.ToLower(to)

	tb := token.BalanceOf(to)
	ts := token.TotalSupply()

	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), tb.Add(tb, amount).String())
	db.Write("TOTAL_SUPPLY", ts.Add(ts, amount).String())
}

func (token *QAN20) transfer(from string, to string, amount *big.Int) bool {
	from = strings.ToLower(from)
	to = strings.ToLower(to)

	fb := token.BalanceOf(from)

	if fb.Cmp(amount) == -1 /* balance is lower than amount */ && fb.Cmp(big.NewInt(0)) != 1 /* balance is lower or equal to 0 */ {
		os.Stderr.WriteString("QAN20: Transfer amount exceeds balance\n")
		os.Exit(1)
	}

	tb := token.BalanceOf(to)

	db.Write(fmt.Sprintf("BALANCE_OF_%v", from), fb.Sub(fb, amount).String())
	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), tb.Add(tb, amount).String())

	return true
}
