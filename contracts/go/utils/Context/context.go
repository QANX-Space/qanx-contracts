package context

import (
	"os"
	"strings"
)

// Retrieves the sender of the transaction
func Sender() string {
	sender := strings.ToLower(os.Getenv("MSG_SENDER"))

	if len(sender) == 0 {
		os.Stderr.WriteString("Context: Sender is not known\n")
		os.Exit(1)
	}

	return sender
}
