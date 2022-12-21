package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
getPosts returns the posts of the user with the given profileUserID.
The posts are returned in the response body.
The response body is a JSON array of posts.
The posts are returned in reverse chronological order.
It's possible to specify the offset and limit of the posts to return.
*/
func (rt *_router) getPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID from the URL
	_profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	profileUserID := uint32(_profileUserID)

	// Check if the user is authorized
	if !isAuthorized(profileUserID, r.Header) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the offset and limit from the query
	limit, offset, err := getLimitAndOffset(r.URL.Query())
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the posts from the database
	posts, err := rt.db.GetPosts(profileUserID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting posts")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		ctx.Logger.WithError(err).Error("Error encoding response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
