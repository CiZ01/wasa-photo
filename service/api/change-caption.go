package api

import (
	"net/http"
	"strconv"

	"encoding/json"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) changeCaption(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the postID from the URL
	postID, err := strconv.Atoi(ps.ByName("postID"))
	if err != nil {
		return
	}

	// Get the profileUserID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		return
	}

	// Get the userID from the context
	userID := ctx.UserID

	// Check if the user is the owner of the post
	if userID != profileUserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Read the request body, save the new caption in the variable NewCaptionBody
	type NewCaptionBody struct {
		NewCaption string `json:"caption"`
	}

	var newCaptionBody NewCaptionBody
	err = json.NewDecoder(r.Body).Decode(&newCaptionBody)
	if err != nil {
		return
	}

	// Update the caption in the database
	err = rt.db.UpdateCaption(userID, postID, newCaptionBody.NewCaption)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
