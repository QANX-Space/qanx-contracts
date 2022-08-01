/*
## Example

```go

package main

import (
	"os"

	qrc721 "qanx.space/qanx-contracts/go/QRC721"
	callintrpr "qanx.space/qanx-contracts/go/utils/CallInterpreter"
)

func main() {
	token := qrc721.NewQRC721("Example", "XMPL", "https://example.com/")

	callintrpr.Interpret(token, os.Args[1:])
}

```*/
package qrc721

import (
	"fmt"
	"os"
	"strconv"

	context "qanx.space/qanx-contracts/go/utils/Context"
	db "qanx.space/qanx-contracts/go/utils/Database"
)

type QRC721Token interface {
	Name() string
	Symbol() string
	BalanceOf(owner string) uint64
	OwnerOf(tokenId uint64) string
	TokenURI(tokenId uint64) string
	Approve(to string, tokendId uint64)
	GetApproved(tokenId uint64) string
	SetApprovalForAll(operator string, approved bool)
	IsApprovedForAll(owner string, operator string) bool
	TransferFrom(from string, to string, tokenId uint64)
	Mint(to string, tokenId uint64)
}

// QRC721 smart contract standard
type QRC721 struct {
	// name of the contract
	name string
	// symbol of the contract
	symbol string
	// prefix for token uris
	baseUri string
}

// Creates the QRC721 smart contract
func NewQRC721(name string, symbol string, baseUri string) *QRC721 {
	return &QRC721{name: name, symbol: symbol, baseUri: baseUri}
}

// Retrieve the name
func (token *QRC721) Name() string {
	return token.name
}

// Retrieve the symbol/ticker
func (token *QRC721) Symbol() string {
	return token.symbol
}

// Retrieve the balance of owner
func (token *QRC721) BalanceOf(owner string) uint64 {
	n, _ := strconv.Atoi(db.Read(fmt.Sprintf("BALANCE_OF_%v", owner)))
	return uint64(n)
}

// Retrieve the owner of token id
func (token *QRC721) OwnerOf(tokenId uint64) string {
	return db.Read(fmt.Sprintf("OWNER_OF_%v", tokenId))
}

// Retrieve the token uri for token id
func (token *QRC721) TokenURI(tokenId uint64) string {
	return token.baseUri + strconv.FormatUint(tokenId, 10)
}

// Give permission to "to" to transfer token id to another account
func (token *QRC721) Approve(to string, tokenId uint64) {
	sender := context.Sender()
	owner := token.OwnerOf(tokenId)

	if to == owner {
		os.Stderr.WriteString("QRC721: Approval to current owner\n")
		os.Exit(1)
	}

	if sender != owner && !token.IsApprovedForAll(owner, sender) {
		os.Stderr.WriteString(fmt.Sprintf("QRC721: %v is not the owner of token id %v\n", sender, tokenId))
		os.Exit(1)
	}

	db.Write(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId), to)
}

// Returns the account approved for token id
func (token *QRC721) GetApproved(tokenId uint64) string {
	return db.Read(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId))
}

// Approve or remove an operator for the caller
func (token *QRC721) SetApprovalForAll(operator string, approved bool) {
	db.Write(fmt.Sprintf("OPERATOR_APPROVAL_%v_%v", context.Sender(), operator), strconv.FormatBool(approved))
}

// Returns if the operator is allowed to manage all of the assets of owner
func (token *QRC721) IsApprovedForAll(owner string, operator string) bool {
	b, _ := strconv.ParseBool(db.Read(fmt.Sprintf("OPERATOR_APPROVAL_%v_%v", owner, operator)))
	return b
}

// Returns if the operator is allowed to manage all of the assets of owner or is the owner
func (token *QRC721) IsApprovedOrOwner(spender string, tokenId uint64) bool {
	owner := token.OwnerOf(tokenId)
	return spender == owner || token.IsApprovedForAll(owner, spender) || token.GetApproved(tokenId) == spender
}

// Transfers token id from "from" to "to"
func (token *QRC721) TransferFrom(from string, to string, tokenId uint64) {
	sender := context.Sender()

	if token.IsApprovedOrOwner(sender, tokenId) {
		os.Stderr.WriteString(fmt.Sprintf("QRC721: %v is not the token owner or approved\n", sender))
		os.Exit(1)
	}

	db.Prune(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId))
	db.Write(fmt.Sprintf("BALANCE_OF_%v", from), strconv.FormatUint(token.BalanceOf(from)-1, 10))
	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), strconv.FormatUint(token.BalanceOf(to)+1, 10))
	db.Write(fmt.Sprintf("OWNER_OF_%v", tokenId), to)
}

// Mints the token id and transfers it to "to"
func (token *QRC721) Mint(to string, tokenId uint64) {
	if len(token.OwnerOf(tokenId)) > 0 {
		os.Stderr.WriteString(fmt.Sprintf("QRC721: Token id %v is already minted\n", tokenId))
		os.Exit(1)
	}

	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), strconv.FormatUint(token.BalanceOf(to)+1, 10))
	db.Write(fmt.Sprintf("OWNER_OF_%v", tokenId), to)
}
