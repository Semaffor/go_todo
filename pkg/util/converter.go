package util

import (
	"fmt"
)

func ConvertToInt(valueForParse string) int {
	var parsedInt int
	_, err := fmt.Sscan(valueForParse, &parsedInt)
	if err != nil {
		return 0
	}
	return parsedInt
}
