package api

import (
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	_profileUserID := ps.ByName("profileUserID")
	profileUserID, err := strconv.Atoi(_profileUserID)
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user is authorized
	userID := isAuthorized(r.Header)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	isBanned, err := rt.db.IsBanned(uint32(profileUserID), userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error checking if user is banned")
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}
	isBanner, err := rt.db.IsBanned(userID, uint32(profileUserID))
	if err != nil {
		ctx.Logger.WithError(err).Error("error checking if user is banner")
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	if isBanned || isBanner {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	_postID := ps.ByName("postID")
	postID, err := strconv.Atoi(_postID)
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	err = rt.db.CreateLike(userID, uint32(profileUserID), uint32(postID))
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: Like.userID, Like.ownerID, Like.postID" {
			http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
			return
		}
		ctx.Logger.WithError(err).Error("error liking post")
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
