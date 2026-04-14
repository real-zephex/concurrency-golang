package scripts

import "strings"

func blockCurl(userAgent string) bool {
	contains := strings.Contains(userAgent, "curl")
	if contains {
		return true
	}
	return false
}
