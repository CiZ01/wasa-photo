package api

import (
	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	isBanned, err := rt.db.IsBanned(profileUserID, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error checking if user is banned")
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}
	isBanner, err := rt.db.IsBanned(userID, profileUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error checking if user is banner")
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	if isBanned || isBanner {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(ps.ByName("postID"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	err = rt.db.CreateLike(userID, profileUserID, postID)
	if err != nil {
		if sqlite3Err, ok := err.(sqlite3.Error); ok {
			if sqlite3Err.Code == sqlite3.ErrConstraint &&
				sqlite3Err.ExtendedCode == 1555 {
				http.Error(w, "Bad Request like already added", http.StatusBadRequest)
				return
			}
		}
		ctx.Logger.WithError(err).Error("error liking post")
		http.Error(w, "Internal Server Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
