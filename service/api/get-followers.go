package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// MISSING LOGGER ERRORS
func (rt *_router) getMyFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and targetUserID from the URL
	_profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	profileUserID := uint32(_profileUserID)

	// Check if the user is authorized
	if !isAuthorized(profileUserID, r.Header) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	limit, offset, err := getLimitAndOffset(r.URL.Query())
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the followers
	dbFollowers, err := rt.db.GetFollowers(profileUserID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while getting the followers")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert the database followers to the API followers
	var followers []User
	for _, dbUser := range dbFollowers {
		var user User
		user.FromDatabase(dbUser)
		followers = append(followers, user)
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(followers); err != nil {
		ctx.Logger.WithError(err).Error("Error while encoding the followers")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
