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
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_userID := ps.ByName("userID")
	userID, err := strconv.Atoi(_userID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Check if the user is authorized
	if !isAuthorized(uint32(userID), r.Header) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	isBanned, err := rt.db.IsBanned(uint32(profileUserID), uint32(userID))
	if err != nil {
		ctx.Logger.WithError(err).Error("error checking if user is banned")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	isBanner, err := rt.db.IsBanned(uint32(userID), uint32(profileUserID))
	if err != nil {
		ctx.Logger.WithError(err).Error("error checking if user is banner")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isBanned || isBanner {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_postID := ps.ByName("postID")
	postID, err := strconv.Atoi(_postID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = rt.db.CreateLike(uint32(userID), uint32(profileUserID), uint32(postID))
	if err != nil {
		ctx.Logger.WithError(err).Error("error liking post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
