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

	// Get the bio from the request body
	type Bio struct {
		Text string `json:"bio"`
	}
	var bio Bio

	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		ctx.Logger.WithError(err).Error("error decoding bio")
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Set the bio
	if err := rt.db.SetBio(profileUserID, bio.Text); err != nil {
		ctx.Logger.WithError(err).Error("error setting bio")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the new bio
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bio); err != nil {
		ctx.Logger.WithError(err).Error("error encoding bio")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
