package api

import (
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and targetUserID from the URL
	_profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	profileUserID := uint32(_profileUserID)

	// Check if the user is authorized
	userID := isAuthorized(r.Header)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if profileUserID == userID {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	isFollowing, err := rt.db.IsFollowing(userID, profileUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while checking if the user is following")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if isFollowing {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	isBanned, err := rt.db.IsBanned(profileUserID, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while checking if the user is banned")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if isBanned {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = rt.db.CreateFollow(userID, profileUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while creating the follow")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
