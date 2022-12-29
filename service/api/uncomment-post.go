package api

import (
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
uncommentPhoto is the handler for the DELETE /profiles/:profileUserID/posts/:postID/comments/:commentID endpoint
It deletes the comment with the given commentID from the post with the given postID.
The comment is deleted only if the user with the given profileUserID is the owner of the comment.
The comment is deleted from the database.
The response body is empty, return a 200 OK status code.
*/

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and photoID from the URL
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

	_postID, err := strconv.Atoi(ps.ByName("postID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	photoID := uint32(_postID)

	// Get commentID from query
	_commentID, err := strconv.Atoi(ps.ByName("commentID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	commentID := uint32(_commentID)

	// Delete the comment
	err = rt.db.DeleteComment(commentID, profileUserID, photoID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error uncommenting photo")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
