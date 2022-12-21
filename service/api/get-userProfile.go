package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the user ID from the URL
	_profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	profileUserID := uint32(_profileUserID)

	if !isAuthorized(profileUserID, r.Header) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the user ID from the header, convert it to uint32.
	// I don't check for errors because the isAuthorized function already did it.
	_userID, _ := strconv.Atoi(r.Header["Authorization"][0])
	userID := uint32(_userID)

	// Check if the user is banned
	isBanned, err := rt.db.IsBanned(profileUserID, userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while checking if the user is banned")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if isBanned {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Get the user from the database
	profile, err := rt.db.GetUserProfile(profileUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while getting the user profile")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		ctx.Logger.WithError(err).Error("Error while encoding the response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
