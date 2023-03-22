package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
commentPhoto is the handler for the POST /users/:profileUserID/posts/:postID/comments endpoint
It creates a new comment for the specified post and returns the new comment
*/
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the user ID from the request
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Invalid profileUserID", http.StatusBadRequest)
		return
	}

	// Get the post ID from the request
	postID, err := strconv.Atoi(ps.ByName("postID"))
	if err != nil {
		http.Error(w, "Invalid postID", http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

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

	// Retrieve the comment from the request body
	var tmpComment Comment

	err = json.NewDecoder(r.Body).Decode(&tmpComment)
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the comment text is valid
	if !tmpComment.isValid() {
		http.Error(w, "Bad reequest invalid comment", http.StatusBadRequest)
		return
	}

	// Create the comment in the database
	dbComment, err := rt.db.CreateComment(userID, profileUserID, postID, tmpComment.Text)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error creating comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert the comment to the API format
	var comment Comment
	err = comment.FromDatabase(dbComment)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error converting comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the comment
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		ctx.Logger.WithError(err).Error("Error encoding comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
