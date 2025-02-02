package api

import (
	"errors"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"git.francescofazzari.it/wasa_photo/service/database"
	"github.com/julienschmidt/httprouter"
)

/*
unlikePhoto is the handler for the DELETE /users/{profileUserID}/posts/{postID}/likes/{userID} endpoint.
It removes a like from the post with the given ID.
*/
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and postID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(ps.ByName("postID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	targetUserID, err := strconv.Atoi(ps.ByName("userID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	if targetUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	isBanned, err := rt.db.IsBanned(profileUserID, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking if user is banned")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if isBanned {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	err = rt.db.DeleteLike(profileUserID, postID, userID)
	if err != nil {
		if errors.Is(err, database.ErrNoRowsAffected) {
			http.Error(w, "Bad Request there is no one like", http.StatusNotFound)
			return
		}
		ctx.Logger.WithError(err).Error("Error unliking photo")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
