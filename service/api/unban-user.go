package api

import (
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and targetUserID from the URL
	_profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	_targetUserID, err := strconv.Atoi(ps.ByName("targetUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	profileUserID := uint32(_profileUserID)
	targetUserID := uint32(_targetUserID)

	// Check if the user is authorized
	if !isAuthorized(profileUserID, r.Header) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if profileUserID == targetUserID {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// I dont know if this is needed
	// isBanned, err := rt.db.IsBanned(targetUserID, profileUserID)
	// if err != nil {
	// 	ctx.Logger.WithError(err).Error("Error checking if the user is banned")
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	// if !isBanned {
	// 	http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
	// 	return
	// }

	err = rt.db.DeleteBan(profileUserID, targetUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting ban")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
