package util

import (
	"strings"
)

func Trim(s string) string {
	trimed := strings.Replace(s, "\n", "", -1)
	trimed = strings.TrimSpace(trimed)
	return trimed
}

func Contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
