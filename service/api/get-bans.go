package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyBans(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Get limit and offset from the queries
	limit, offset, err := getLimitAndOffset(r.URL.Query())
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	bans, err := rt.db.GetBans(profileUserID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting bans")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
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
