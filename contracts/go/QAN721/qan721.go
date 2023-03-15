/*
## Example

```go

package main

import (
	"os"

	qan721 "qanx.space/qanx-contracts/go/QAN721"
	callintrpr "qanx.space/qanx-contracts/go/utils/CallInterpreter"
)

func main() {
	token := qan721.NewQAN721("Example", "XMPL", "https://example.com/")

	callintrpr.Interpret(token, os.Args[1:])
}

```*/
package qan721

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	common "qanx.space/qanx-contracts/go/utils/Common"
	db "qanx.space/qanx-contracts/go/utils/Database"
	message "qanx.space/qanx-contracts/go/utils/Message"
)

type QAN721Token interface {
	Name() string
	Symbol() string
	BalanceOf(owner string) *big.Int
	OwnerOf(tokenId *big.Int) string
	TokenURI(tokenId *big.Int) string
	Approve(to string, tokendId *big.Int)
	GetApproved(tokenId *big.Int) string
	SetApprovalForAll(operator string, approved bool)
	IsApprovedForAll(owner string, operator string) bool
	TransferFrom(from string, to string, tokenId *big.Int)
	Mint(to string, tokenId *big.Int)
}

// QAN721 smart contract standard
type QAN721 struct {
	// name of the contract
	name string
	// symbol of the contract
	symbol string
	// prefix for token uris
	baseUri string
}

// Creates the QAN721 smart contract
func NewQAN721(name string, symbol string, baseUri string) *QAN721 {
	return &QAN721{name: name, symbol: symbol, baseUri: baseUri}
}

// Retrieve the name
func (token *QAN721) Name() string {
	return token.name
}

// Retrieve the symbol/ticker
func (token *QAN721) Symbol() string {
	return token.symbol
}

// Retrieve the balance of owner
func (token *QAN721) BalanceOf(owner string) *big.Int {
	owner = strings.ToLower(owner)

	n, _ := common.ParseBig256(db.Read(fmt.Sprintf("BALANCE_OF_%v", owner)))
	return n
}

// Retrieve the owner of token id
func (token *QAN721) OwnerOf(tokenId *big.Int) string {
	return db.Read(fmt.Sprintf("OWNER_OF_%v", tokenId.String()))
}

// Retrieve the token uri for token id
func (token *QAN721) TokenURI(tokenId *big.Int) string {
	return token.baseUri + tokenId.String()
}

// Give permission to "to" to transfer token id to another account
func (token *QAN721) Approve(to string, tokenId *big.Int) {
	to = strings.ToLower(to)
	sender := message.Sender()
	owner := token.OwnerOf(tokenId)

	if to == owner {
		os.Stderr.WriteString("QAN721: Approval to current owner\n")
		os.Exit(1)
	}

	if sender != owner && !token.IsApprovedForAll(owner, sender) {
		os.Stderr.WriteString(fmt.Sprintf("QAN721: %v is not the owner of token id %v\n", sender, tokenId.String()))
		os.Exit(1)
	}

	db.Write(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId.String()), to)
}

// Returns the account approved for token id
func (token *QAN721) GetApproved(tokenId *big.Int) string {
	return db.Read(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId.String()))
}

// Approve or remove an operator for the caller
func (token *QAN721) SetApprovalForAll(operator string, approved bool) {
	operator = strings.ToLower(operator)

	db.Write(fmt.Sprintf("OPERATOR_APPROVAL_%v_%v", message.Sender(), operator), strconv.FormatBool(approved))
}

// Returns if the operator is allowed to manage all of the assets of owner
func (token *QAN721) IsApprovedForAll(owner string, operator string) bool {
	owner = strings.ToLower(owner)
	operator = strings.ToLower(operator)

	b, _ := strconv.ParseBool(db.Read(fmt.Sprintf("OPERATOR_APPROVAL_%v_%v", owner, operator)))
	return b
}

// Returns if the operator is allowed to manage all of the assets of owner or is the owner
func (token *QAN721) IsApprovedOrOwner(spender string, tokenId *big.Int) bool {
	spender = strings.ToLower(spender)
	owner := token.OwnerOf(tokenId)

	return spender == owner || token.IsApprovedForAll(owner, spender) || token.GetApproved(tokenId) == spender
}

// Transfers token id from "from" to "to"
func (token *QAN721) TransferFrom(from string, to string, tokenId *big.Int) {
	from = strings.ToLower(from)
	to = strings.ToLower(to)
	sender := message.Sender()

	if !token.IsApprovedOrOwner(sender, tokenId) {
		os.Stderr.WriteString(fmt.Sprintf("QAN721: %v is not the token owner or approved\n", sender))
		os.Exit(1)
	}

	amount := big.NewInt(1)

	fb := token.BalanceOf(from)
	tb := token.BalanceOf(to)

	db.Write(fmt.Sprintf("BALANCE_OF_%v", from), fb.Sub(fb, amount).Text(16))
	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), tb.Add(tb, amount).Text(16))

	db.Prune(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId.String()))
	db.Write(fmt.Sprintf("OWNER_OF_%v", tokenId.String()), to)
}

// Mints the token id and transfers it to "to"
func (token *QAN721) Mint(to string, tokenId *big.Int) {
	to = strings.ToLower(to)

	if len(token.OwnerOf(tokenId)) > 0 {
		os.Stderr.WriteString(fmt.Sprintf("QAN721: Token id %v is already minted\n", tokenId.String()))
		os.Exit(1)
	}

	amount := big.NewInt(1)

	tb := token.BalanceOf(to)

	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), tb.Add(tb, amount).Text(16))
	db.Write(fmt.Sprintf("OWNER_OF_%v", tokenId.String()), to)
}
