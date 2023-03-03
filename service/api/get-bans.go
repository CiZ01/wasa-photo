package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"git.francescofazzari.it/wasa_photo/service/api/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyBans(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and targetUserID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	// Check if the user is authorized
	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Get limit and offset from the queries
	limit, offset, err := utils.GetLimitAndOffset(r.URL.Query())
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	dbBans, err := rt.db.GetBans(profileUserID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting bans")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	bans := make([]User, len(dbBans))

	for i, dbBan := range dbBans {
		var user User
		err := user.FromDatabase(dbBan)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error while converting the user")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		bans[i] = user
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bans); err != nil {
		ctx.Logger.WithError(err).Error("Error encoding bans")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
