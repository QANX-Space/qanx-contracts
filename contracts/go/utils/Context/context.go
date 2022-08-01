package context

import "os"

// Retrieves the sender of the transaction
func Sender() string {
	sender := os.Getenv("SENDER")

	if len(sender) == 0 {
		os.Stderr.WriteString("Context: Sender is not known\n")
		os.Exit(1)
	}

	return sender
}
