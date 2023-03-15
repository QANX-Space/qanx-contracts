package common

import "math/big"

// Taken from the go-ethereum repo
// https://github.com/ethereum/go-ethereum/blob/f86913bc3e9a4f2439b6c3cd4d00cb364495238c/common/math/big.go#L117-L132
func ParseBig256(s string) (*big.Int, bool) {
	if s == "" {
		return new(big.Int), true
	}
	var bigint *big.Int
	var ok bool
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		bigint, ok = new(big.Int).SetString(s[2:], 16)
	} else {
		bigint, ok = new(big.Int).SetString(s, 10)
	}
	if ok && bigint.BitLen() > 256 {
		bigint, ok = nil, false
	}
	return bigint, ok
}
