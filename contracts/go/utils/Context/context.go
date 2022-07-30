package context

import "os"

func Sender() string {
	return os.Getenv("SENDER")
}
