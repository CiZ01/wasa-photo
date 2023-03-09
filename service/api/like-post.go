package api

import (
	"errors"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"github.com/mattn/go-sqlite3"
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

	if isBanned {
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
		var sqlite3Err sqlite3.Error
		if errors.As(err, &sqlite3Err) {
			if sqlite3Err.Code == sqlite3.ErrConstraint && sqlite3Err.ExtendedCode == 1555 {
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
