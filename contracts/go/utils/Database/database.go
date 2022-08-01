package db

import (
	"fmt"
	"os"
)

// Write key and value to the database
func Write(key string, value string) {
	os.Stdout.WriteString(fmt.Sprintf("DBW=%v=%v\n", key, value))
	os.Setenv(fmt.Sprintf("DB_%v", key), value)
}

// Removes the key from the database
func Prune(key string) {
	os.Stdout.WriteString(fmt.Sprintf("DBP=%v\n", key))
	os.Setenv(fmt.Sprintf("DB_%v", key), "")
}

// Read key from database
func Read(key string) string {
	value := os.Getenv(fmt.Sprintf("DB_%v", key))

	if len(value) == 0 {
		os.Stderr.WriteString(fmt.Sprintf("Database: Can't find key \"DB_%v\" in env\n", key))
		os.Exit(1)
	}

	return value
}
