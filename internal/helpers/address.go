package helpers

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Address is a wrapper around common.Address with additional functionalities
type Address struct {
	common.Address
}

// Scan implements Scanner.Scan for common.Address
func (a *Address) Scan(value interface{}) error {
	hex := strings.TrimPrefix(value.(string), "0x")
	a.SetBytes(common.Hex2Bytes(hex))
	return nil
}
