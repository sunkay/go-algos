// --- Directions
// Given an integer, return an integer that is the reverse
// ordering of numbers.
// --- Examples
//   reverseInt(15) === 51
//   reverseInt(981) === 189
//   reverseInt(500) === 5
//   reverseInt(-15) === -51
//   reverseInt(-90) === -9

package main

import (
	"strconv"
	"strings"
)

func reverseInt(v int) int {
	// handle zero
	if v == 0 {
		return v
	}

	isNegative := (v < 0)

	valueStr := strconv.Itoa(v)

	// trim trailing zero's
	valueStr = strings.Trim(valueStr, "0-")

	// reverse string
	revValueStr := ""
	for i := 0; i < len(valueStr); i++ {
		revValueStr = string(valueStr[i]) + revValueStr
	}

	revValue, _ := strconv.Atoi(revValueStr)

	if isNegative {
		revValue = -revValue
	}

	return revValue
}
