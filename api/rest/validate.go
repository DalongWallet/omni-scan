package rest

import "strconv"

func isInt(param string) bool {
	if _, err := strconv.Atoi(param); err != nil {
		return false
	}
	return true
}

