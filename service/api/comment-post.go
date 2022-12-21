package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the user ID from the request
	_profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Invalid profileUserID", http.StatusBadRequest)
		return
	}
	profileUserID := uint32(_profileUserID)

	// Get the post ID from the request
	_postID, err := strconv.Atoi(ps.ByName("postID"))
	if err != nil {
		http.Error(w, "Invalid postID", http.StatusBadRequest)
		return
	}
	postID := uint32(_postID)

	// Get the user ID from the request
	_userID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}
	userID := uint32(_userID)

	if !isAuthorized(uint32(_userID), r.Header) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	isBanned, err := rt.db.IsBanned(profileUserID, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while checking if the user is banned")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if isBanned {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var tmpComment Comment

	err = json.NewDecoder(r.Body).Decode(&tmpComment)
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	comment, err := rt.db.CreateComment(userID, profileUserID, postID, tmpComment.Text)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error creating comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}
