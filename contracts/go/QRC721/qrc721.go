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

type QRC721 struct {
	name    string
	symbol  string
	baseUri string
}

func NewQRC721(name string, symbol string, baseUri string) *QRC721 {
	return &QRC721{name: name, symbol: symbol, baseUri: baseUri}
}

func (token *QRC721) Name() string {
	return token.name
}

func (token *QRC721) Symbol() string {
	return token.symbol
}

func (token *QRC721) BalanceOf(owner string) uint64 {
	n, _ := strconv.Atoi(db.Read(fmt.Sprintf("BALANCE_OF_%v", owner)))
	return uint64(n)
}

func (token *QRC721) OwnerOf(tokenId uint64) string {
	return db.Read(fmt.Sprintf("OWNER_OF_%v", tokenId))
}

func (token *QRC721) TokenURI(tokenId uint64) string {
	return token.baseUri + strconv.FormatUint(tokenId, 10)
}

func (token *QRC721) Approve(to string, tokenId uint64) {
	sender := context.Sender()

	if token.OwnerOf(tokenId) != sender {
		os.Stderr.WriteString(fmt.Sprintf("QRC721: %v is not the owner of %v\n", sender, tokenId))
		os.Exit(1)
	}

	db.Write(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId), to)
}

func (token *QRC721) GetApproved(tokenId uint64) string {
	return db.Read(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId))
}

func (token *QRC721) SetApprovalForAll(operator string, approved bool) {
	db.Write(fmt.Sprintf("OPERATOR_APPROVAL_%v_%v", context.Sender(), operator), strconv.FormatBool(approved))
}

func (token *QRC721) IsApprovedForAll(owner string, operator string) bool {
	b, _ := strconv.ParseBool(db.Read(fmt.Sprintf("OPERATOR_APPROVAL_%v_%v", owner, operator)))
	return b
}

func (token *QRC721) TransferFrom(from string, to string, tokenId uint64) {
	if token.OwnerOf(tokenId) != from {
		os.Stderr.WriteString(fmt.Sprintf("QRC721: %v is not the owner of token id %v\n", from, tokenId))
		os.Exit(1)
	}

	db.Prune(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId))
	db.Write(fmt.Sprintf("BALANCE_OF_%v", from), strconv.FormatUint(token.BalanceOf(from)-1, 10))
	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), strconv.FormatUint(token.BalanceOf(to)+1, 10))
	db.Write(fmt.Sprintf("OWNER_OF_%v", tokenId), to)
}

func (token *QRC721) Mint(to string, tokenId uint64) {
	if len(token.OwnerOf(tokenId)) > 0 {
		os.Stderr.WriteString(fmt.Sprintf("QRC721: Token id %v is already minted\n", tokenId))
		os.Exit(1)
	}

	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), strconv.FormatUint(token.BalanceOf(to)+1, 10))
	db.Write(fmt.Sprintf("OWNER_OF_%v", tokenId), to)
}
