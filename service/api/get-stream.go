package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/utils"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
getMyStream is the handler for the GET /users/:profileUserID/feed endpoint.
It returns the posts of the followed users of the user with the profileUserID.
*/
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Get limit and offset from the queries
	limit, offset, err := utils.GetLimitAndOffset(r.URL.Query())
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the stream
	dbStream, err := rt.db.GetStream(profileUserID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting the stream")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var posts []Post

	// Convert the posts from database.Post to api.Post
	for _, dbPost := range dbStream {
		var post Post
		err = post.FromDatabase(dbPost)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error converting post")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		ctx.Logger.WithError(err).Error("Error encoding the stream")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
