package util

import (
	"strconv"
)

// ToInt 转化为整数。
func ToInt(str string, defaultValue int) int {
	n, err := strconv.Atoi(str)
	if err == nil {
		return n
	}

	return defaultValue
}
