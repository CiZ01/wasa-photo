package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyBio(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID from the URL
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

	// Get the bio from the request body
	var bio string
	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Set the bio
	if err := rt.db.SetBio(profileUserID, bio); err != nil {
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the new bio
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bio); err != nil {
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
		return
	}
}
