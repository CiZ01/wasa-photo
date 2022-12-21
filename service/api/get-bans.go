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
	if !isAuthorized(profileUserID, r.Header) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get limit and offset from the queries
	limit, offset, err := getLimitAndOffset(r.URL.Query())
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	bans, err := rt.db.GetBans(profileUserID, limit, offset)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting bans")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(bans)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error encoding bans")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
