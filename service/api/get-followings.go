package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/utils"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// MISSING LOGGER ERRORS

/*
getMyFollowings is the handler for GET /users/:profileUserID/followings
Return the followings of the user with the given profileUserID.
Only the user with the given profileUserID can access this endpoint.
*/
func (rt *_router) getMyFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and targetUserID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user is authorized
	if profileUserID != ctx.UserID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	limit, offset, err := utils.GetLimitAndOffset(r.URL.Query())
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the followings
	dbFollowings, err := rt.db.GetFollowings(profileUserID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while getting the followings")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert the database followings to the API followings
	followings := make([]User, len(dbFollowings))
	for i, dbUser := range dbFollowings {
		var user User
		err := user.FromDatabase(dbUser)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error while converting the user")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		followings[i] = user
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(followings); err != nil {
		ctx.Logger.WithError(err).Error("Error while encoding the response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
