package main

import (
	"fmt"

	qrc721 "qanx.space/qanx-contracts/go/QRC721"
)

func main() {
	token := qrc721.QRC721{
		name:    "Example Token",
		symbol:  "EXPL",
		baseUri: "https://example.com/",
	}

	fmt.Println(token)
}
