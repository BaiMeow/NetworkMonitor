package utils

import (
	"fmt"
	"math"
	"strconv"
)

func wrapASNErr(value any) error {
	return fmt.Errorf("invalid asn: %v", value)
}

func MustASN(input any) (uint32, error) {
	switch asn := input.(type) {
	case int:
		if asn > math.MaxUint32 || asn < 0 {
			return 0, wrapASNErr(asn)
		}
		return uint32(asn), nil
	case string:
		uintASN, err := strconv.ParseUint(asn, 10, 32)
		if err != nil {
			return 0, wrapASNErr(asn)
		}
		return uint32(uintASN), nil
	default:
		return 0, wrapASNErr(input)
	}
}
