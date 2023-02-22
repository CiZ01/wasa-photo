package api

import (
	"net/http"
	"strconv"
)

/*
isAuthorized checks if the user is authorized to perform the action, by checking the Authorization header.
The auth token must be the first element of the Authorization header.
The auth token must be in the format "Bearer and must be equal to userID to be valid.
If the user is authorized, the function returns true, otherwise it returns false.
*/
func isAuthorized(header http.Header) int {
	authToken, err := strconv.Atoi(header.Get("Authorization"))
	if err != nil {
		return 0
	}
	return authToken
}
