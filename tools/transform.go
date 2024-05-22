package tools

import (
	"strconv"
)

func UintToStr(num uint) string {
	return strconv.Itoa(int(num))
}

func StrToUint(str string) uint {
	if num, err := strconv.ParseUint(str, 10, 64); err == nil {
		return uint(num)
	}
	return 0
}
