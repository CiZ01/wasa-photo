package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and photoID from the URL
	_profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	profileUserID := uint32(_profileUserID)

	_postID, err := strconv.Atoi(ps.ByName("postID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	photoID := uint32(_postID)

	// Check if the user is authorized
	userID := isAuthorized(r.Header)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
	limit, offset, err := getLimitAndOffset(r.URL.Query())
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	likes, err := rt.db.GetLikes(photoID, profileUserID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting likes")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(likes); err != nil {
		ctx.Logger.WithError(err).Error("Error encoding likes")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
