package api

import (
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
* FollowUser is the handler for the PUT /profiles/:profileUserID/followings/:targetUser API
* It allows a user to follow another user
* It returns 400 if the user is trying to follow himself
* It returns 400 if the user is already following the target user
* It returns 400 if the target user has banned the user
* It returns 500 if there is an error while checking if the user is already following the target user

DA RIVEDERE
*/
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and targetUserID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	targetUserID, err := strconv.Atoi(ps.ByName("targetUserID"))

	if profileUserID == targetUserID {
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

	w.WriteHeader(http.StatusOK)
}
