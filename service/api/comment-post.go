package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

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

	dbComment, err := rt.db.CreateComment(userID, profileUserID, postID, tmpComment.Text)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error creating comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var comment Comment
	err = comment.FromDatabase(dbComment)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error converting comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		ctx.Logger.WithError(err).Error("Error encoding comment")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
