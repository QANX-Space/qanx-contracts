package message

import (
	"math/big"
	"os"
	"strings"

	common "qanx.space/qanx-contracts/go/utils/Common"
)

// Retrieves the sender of the transaction
func Sender() string {
	sender := strings.ToLower(os.Getenv("MSG_SENDER"))

	if len(sender) == 0 {
		os.Stderr.WriteString("Message: Sender is not known\n")
		os.Exit(1)
	}

	return sender
}

// Retrieves the QANX value of the transaction
func Value() *big.Int {
	n, _ := common.ParseBig256(os.Getenv("MSG_VALUE"))

	return n
}
