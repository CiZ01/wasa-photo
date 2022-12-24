package api

import (
	"net/url"
	"strconv"
)

/*
getLimitAndOffset returns the limit and offset from the queries.
If the queries are not present or not valid, it returns the default values.
The default values are 10 for limit and 0 for offset.
*/
func getLimitAndOffset(query url.Values) (int32, uint32, error) {
	// Get limit and offset from the queries
	limit, offset := int32(10), uint32(0)

	// Check if the offset is valid
	if query.Has("offset") {
		_offset, err := strconv.Atoi(query.Get("offset"))
		if err != nil {
			return 0, 0, err
		}
		offset = uint32(_offset)
	}

	// Check if the limit is valid
	if query.Has("limit") {
		_limit, err := strconv.Atoi(query.Get("limit"))
		if err != nil {
			return 0, 0, err
		}
		limit = int32(_limit)
	}
	return limit, offset, nil
}
