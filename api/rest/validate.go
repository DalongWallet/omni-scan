package rest

import "strconv"

func isUintStr(param string) bool {
	if n, err := strconv.Atoi(param); err == nil && n >= 0 {
		return true
	}
	return false
}
