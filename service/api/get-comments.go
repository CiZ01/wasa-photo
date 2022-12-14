package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID from the URL
	_profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	profileUserID := uint32(_profileUserID)

	// Get the postID from the URL
	_postID, err := strconv.Atoi(ps.ByName("postID"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	postID := uint32(_postID)

	// Get the userID from the Authorization header
	_userID := r.Header.Get("Authorization")

	userID, err := strconv.Atoi(_userID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if _userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Check if the user is banned
	isBanned, err := rt.db.IsBanned(profileUserID, uint32(userID))
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking if user is banned")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if isBanned {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Get limit and offset from the queries
	offset, limit, err := getLimitAndOffset(ps)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Get the comments
	comments, err := rt.db.GetComments(profileUserID, postID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting comments")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the comments
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error encoding comments")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
