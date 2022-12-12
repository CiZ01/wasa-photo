package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID from the URL
	_profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	profileUserID := uint32(_profileUserID)

	// Check if the user is authorized
	if !isAuthorized(profileUserID, r.Header) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	offset, limit := uint32(0), uint32(10)
	if ps.ByName("offset") != "" {
		_offset, err := strconv.Atoi(ps.ByName("offset"))
		if err != nil {
			http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
			return
		}
		offset = uint32(_offset)
	}

	if ps.ByName("limit") != "" {
		_limit, err := strconv.Atoi(ps.ByName("limit"))
		if err != nil {
			http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
			return
		}
		limit = uint32(_limit)
	}

	// Get the posts from the database
	posts, err := rt.db.GetPosts(profileUserID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting posts")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error encoding response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
