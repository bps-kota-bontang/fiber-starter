package utils

import (
	"strconv"
)

// Convert string ID to uint safely
func ParseID(id string) (uint, error) {
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(uintID), nil
}
