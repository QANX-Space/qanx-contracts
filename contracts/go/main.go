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
