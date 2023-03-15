package db

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func formatSha256(key string) string {
	return fmt.Sprintf("0x%x", sha256.Sum256([]byte(key)))
}

// Write key and value to the database
func Write(key string, value string) {
	sha256key := formatSha256(key)

	os.Stdout.WriteString(fmt.Sprintf("DBW=%v=%v\n", sha256key, value))
	os.Setenv(fmt.Sprintf("DB_%v", sha256key), value)
}

// Removes the key from the database
func Prune(key string) {
	sha256key := formatSha256(key)

	os.Stdout.WriteString(fmt.Sprintf("DBP=%v\n", sha256key))
	os.Setenv(fmt.Sprintf("DB_%v", sha256key), "")
}

// Read key from database
func Read(key string) string {
	sha256key := formatSha256(key)

	value := os.Getenv(fmt.Sprintf("DB_%v", sha256key))

	if len(value) == 0 {
		os.Stderr.WriteString(fmt.Sprintf("Database: Can't find key \"DB_%v\" as \"DB_%v\" in env\n", key, sha256key))
		os.Exit(1)
	}

	return value
}
