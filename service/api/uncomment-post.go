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
*/

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID and postID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(ps.ByName("postID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user is allowed to delete the comment
	if profileUserID != ctx.UserID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Get commentID from query
	commentID, err := strconv.Atoi(ps.ByName("commentID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Delete the comment
	err = rt.db.DeleteComment(commentID, profileUserID, postID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error uncommenting photo")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return a 200 OK status code
	w.WriteHeader(http.StatusOK)
}
