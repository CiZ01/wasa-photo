package utils

import (
	"net/url"
	"strconv"
)

/*
getLimitAndOffset returns the limit and offset from the queries.
If the queries are not present or not valid, it returns the default values.
The default values are 10 for limit and 0 for offset.
*/
func GetLimitAndOffset(query url.Values) (int, int, error) {
	// Get limit and offset from the queries
	limit, offset := 10, 0

	var err error
	// Check if the offset is valid
	if query.Has("offset") {
		offset, err = strconv.Atoi(query.Get("offset"))
		if err != nil {
			return 0, 0, err
		}
	}

	// Check if the limit is valid
	if query.Has("limit") {
		limit, err = strconv.Atoi(query.Get("limit"))
		if err != nil {
			return 0, 0, err
		}
	}

	if limit > 100 {
		limit = 100
	}
	if limit < 0 {
		limit = 0
	}
	if offset < 0 {
		offset = 0
	}

	return limit, offset, nil
}
