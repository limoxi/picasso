package util

import "strconv"

func WrappedInt(num int) string{
	s := strconv.Itoa(num)
	if num < 10{
		s = "0" + s
	}
	return s
}