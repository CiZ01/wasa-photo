package api

import "strconv"

func isAuthorized(userID uint32, header []string) bool {
	if len(header) == 0 {
		return false
	}
	autToken, _ := strconv.Atoi(header[0])
	return uint32(autToken) == userID
}
