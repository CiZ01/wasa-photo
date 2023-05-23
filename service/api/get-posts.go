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
getPosts is the handler for the GET /users/:profileUserID/posts endpoint
It return the posts of the user with the given profileUserID.
*/
func (rt *_router) getPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the offset and limit from the query
	limit, offset, err := utils.GetLimitAndOffset(r.URL.Query())
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID

	isBanned, err := rt.db.IsBanned(userID, profileUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking if the user is banned")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if isBanned {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Get the posts from the database
	dbPosts, err := rt.db.GetPosts(userID, profileUserID, offset, limit)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting posts")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	posts := make([]User, len(dbPosts))

	for i, dbBan := range dbPosts {
		var post Post
		err := post.FromDatabase(dbPost)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error while converting the post")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		posts[i] = post
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		ctx.Logger.WithError(err).Error("Error encoding response")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
