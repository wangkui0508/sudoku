package types

import (
	"math/bits"
)

type MaskType = uint16

func IsFixed(n MaskType) bool {
	return bits.OnesCount16(n) == 1
}

func PossibleValueCount(n MaskType) int {
	return bits.OnesCount16(n)
}

func GetResult(n MaskType) int {
	return bits.TrailingZeros16(n)+1
}

