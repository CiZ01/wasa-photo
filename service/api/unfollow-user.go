package api

import (
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
unfollowUser unfollows the user “ followed by .
*/
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and targetUserID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	targetUserID, err := strconv.Atoi(ps.ByName("targetUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	if profileUserID != ctx.UserID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if profileUserID == targetUserID {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	isBanned, err := rt.db.IsBanned(profileUserID, targetUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking if the user is banned")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if isBanned {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	isFollowing, err := rt.db.IsFollowing(profileUserID, targetUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error checking if the user is following the target user")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !isFollowing {
		http.Error(w, "Bad Request the user is not followed", http.StatusBadRequest)
		return
	}

	err = rt.db.DeleteFollow(profileUserID, targetUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting follow")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
