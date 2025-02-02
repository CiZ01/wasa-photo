package api

import (
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
UnbanUser is the handler for the DELETE /users/{profileUserID}/bans/{targetUserID} endpoint
It deletes the ban between the profileUserID and the targetUserID
*/
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and targetUserID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	targetUserID, err := strconv.Atoi(ps.ByName("targetUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user is trying to unban himself
	if targetUserID == userID {
		http.Error(w, "Bad Request the user trying to unban himself", http.StatusBadRequest)
		return
	}

	// Check if the user is authorized
	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusBadRequest)
		return
	}

	// Delete the ban in the database
	err = rt.db.DeleteBan(profileUserID, targetUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting ban")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
