package api

import (
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and targetUserID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	if profileUserID == userID {
		http.Error(w, "Bad Request", http.StatusBadRequest)
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

	err = rt.db.DeleteBan(userID, profileUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting ban")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
