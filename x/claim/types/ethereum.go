package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// isValidEthAddress checks if the provided string is a valid address or not.
func IsValidEthAddress(address string) bool {
	return common.IsHexAddress(address)
}
