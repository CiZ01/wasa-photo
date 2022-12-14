package api

import (
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func getLimitAndOffset(ps httprouter.Params) (uint32, uint32, error) {
	// Get limit and offset from the queries
	offset, limit := uint32(0), uint32(10)
	if ps.ByName("offset") != "" {
		_offset, err := strconv.Atoi(ps.ByName("offset"))
		if err != nil {
			return 0, 0, err
		}
		offset = uint32(_offset)
	}

	if ps.ByName("limit") != "" {
		_limit, err := strconv.Atoi(ps.ByName("limit"))
		if err != nil {
			return 0, 0, err
		}
		limit = uint32(_limit)
	}
	return offset, limit, nil
}
