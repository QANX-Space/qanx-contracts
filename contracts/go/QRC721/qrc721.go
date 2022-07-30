package qrc721

import (
	"fmt"
	"os"
	"strconv"

	context "qanx.space/qanx-contracts/go/utils/Context"
	db "qanx.space/qanx-contracts/go/utils/Database"
)

type QRC721Token interface {
	balanceOf(owner string) uint64
	ownerOf(tokenId uint64) string
	tokenURI(tokenId uint64) string
	approve(to string, tokendId uint64)
	getApproved(tokenId uint64) string
	setApprovalForAll(operator string, approved bool)
	isApprovedForAll(owner string, operator string) bool
	transferFrom(from string, to string, tokenId uint64)
	mint(to string, tokenId uint64)
}

type QRC721 struct {
	name, symbol, baseUri string
}

func (token QRC721) balanceOf(owner string) uint64 {
	n, _ := strconv.Atoi(db.Read(fmt.Sprintf("BALANCE_OF_%v", owner)))
	return uint64(n)
}

func (token QRC721) ownerOf(tokenId uint64) string {
	return db.Read(fmt.Sprintf("OWNER_OF_%v", tokenId))
}

func (token QRC721) tokenURI(tokenId uint64) string {
	return token.baseUri + db.Read(fmt.Sprintf("TOKEN_URI_%v", tokenId))
}

func (token QRC721) approve(to string, tokenId uint64) {
	sender := context.Sender()

	if token.ownerOf(tokenId) != sender {
		os.Stderr.WriteString(fmt.Sprintf("QRC721: %v is not the owner of %v\n", sender, tokenId))
		return
	}

	db.Write(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId), to)
}

func (token QRC721) getApproved(tokenId uint64) string {
	return db.Read(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId))
}

func (token QRC721) setApprovalForAll(operator string, approved bool) {
	db.Write(fmt.Sprintf("OPERATOR_APPROVAL_%v_%v", context.Sender(), operator), strconv.FormatBool(approved))
}

func (token QRC721) isApprovedForAll(owner string, operator string) bool {
	b, _ := strconv.ParseBool(db.Read(fmt.Sprintf("OPERATOR_APPROVAL_%v_%v", owner, operator)))
	return b
}

func (token QRC721) transferFrom(from string, to string, tokenId uint64) {
	if token.ownerOf(tokenId) != from {
		os.Stderr.WriteString(fmt.Sprintf("QRC721: %v is not the owner of %v\n", from, tokenId))
		return
	}

	db.Prune(fmt.Sprintf("TOKEN_APPROVAL_%v", tokenId))
	db.Write(fmt.Sprintf("BALANCE_OF_%v", from), strconv.FormatUint(token.balanceOf(from)-1, 10))
	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), strconv.FormatUint(token.balanceOf(to)+1, 10))
	db.Write(fmt.Sprintf("OWNER_OF_%v", tokenId), to)
}

func (token QRC721) mint(to string, tokenId uint64) {
	if token.ownerOf(tokenId) != "" {
		os.Stderr.WriteString(fmt.Sprintf("QRC721: %v is already minted\n", tokenId))
		return
	}

	db.Write(fmt.Sprintf("BALANCE_OF_%v", to), strconv.FormatUint(token.balanceOf(to)+1, 10))
	db.Write(fmt.Sprintf("OWNER_OF_%v", tokenId), to)
}
