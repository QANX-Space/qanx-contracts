package db

import (
	"fmt"
	"os"
)

func Write(key string, value string) {
	os.Stdout.WriteString(fmt.Sprintf("DBW=%v=%v\n", key, value))
	os.Setenv(fmt.Sprintf("DB_%v", key), value)
}

func Prune(key string) {
	os.Stdout.WriteString(fmt.Sprintf("DBP=%v\n", key))
	os.Setenv(fmt.Sprintf("DB_%v", key), "")
}

func Read(key string) string {
	return os.Getenv(fmt.Sprintf("DB_%v", key))
}
