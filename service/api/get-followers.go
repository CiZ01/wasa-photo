package api

import (
	"encoding/json"
	"git.francescofazzari.it/wasa_photo/service/api/utils"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
GetMyFollowers is the handler for the GET /users/:profileUserID/followers endpoint
It return the list of the followers of the user with the given profileUserID
*/
func (rt *_router) getMyFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and targetUserID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	limit, offset, err := utils.GetLimitAndOffset(r.URL.Query())
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
	followers := make([]User, len(dbFollowers))
	for i, dbUser := range dbFollowers {
		var user User
		err := user.FromDatabase(dbUser)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error while converting the user")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		followers[i] = user
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
