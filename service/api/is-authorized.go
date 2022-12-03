package api

import (
	"net/http"
	"strconv"
	"strings"
)

/*
isAuthorized checks if the user is authorized to perform the action, by checking the Authorization header.
The auth token must be the first element of the Authorization header.
The auth token must be in the format "Bearer and must be equal to userID to be valid.
If the user is authorized, the function returns true, otherwise it returns false.
*/
func isAuthorized(userID uint32, header http.Header) bool {
	headerToList := strings.Split(header.Get("Authorization"), "")

	if len(headerToList) == 0 {
		return false
	}
	autToken, _ := strconv.Atoi(headerToList[0])
	return uint32(autToken) == userID
}
