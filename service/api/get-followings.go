package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// MISSING LOGGER ERRORS
func (rt *_router) getMyFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	} else if userID != profileUserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	limit, offset, err := getLimitAndOffset(r.URL.Query())
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the followings
	followings, err := rt.db.GetFollowings(profileUserID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while getting the followings")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(followings); err != nil {
		ctx.Logger.WithError(err).Error("Error while encoding the response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
