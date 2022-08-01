package context

import "os"

// Retrieves the sender of the transaction
func Sender() string {
	return os.Getenv("SENDER")
}
