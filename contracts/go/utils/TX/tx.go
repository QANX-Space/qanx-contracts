package tx

import (
	"math/big"
	"os"
	"strings"

	common "qanx.space/qanx-contracts/go/utils/Common"
)

// Retrieves the gas price of the transaction
func GasPrice() *big.Int {
	n, _ := common.ParseBig256(os.Getenv("TX_GASPRICE"))

	return n
}

// Retrieves the origin address of the transaction
func Origin() string {
	origin := strings.ToLower(os.Getenv("TX_ORIGIN"))

	if len(origin) == 0 {
		os.Stderr.WriteString("TX: Origin is not known\n")
		os.Exit(1)
	}

	return origin
}
