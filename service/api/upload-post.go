package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) UploadPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	_profileUserID := ps.ByName("profileUserID")
	profileUserID, err := strconv.Atoi(_profileUserID)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Check if the user is authorized
	if !isAuthorized(uint32(profileUserID), r.Header) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		ctx.Logger.WithError(err).Error("error parsing multipart form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	file, _, err := r.FormFile("image")
	if err != nil {
		ctx.Logger.WithError(err).Error("error getting photo")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer file.Close()

	dbuser, err := rt.db.GetUserByID(uint32(profileUserID))
	if err != nil {
		ctx.Logger.WithError(err).Error("error getting user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user User
	user.FromDatabase(dbuser)

	// Create the post
	var newPost = Post{
		User:    user,
		Caption: string(r.FormValue("caption")),
	}

	dbPost := newPost.ToDatabase(user)

	dbNewPost, err := rt.db.CreatePost(dbPost)
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newPost = newPost.FromDatabase(dbNewPost, dbuser)

	tmpfile, err := os.Create(newPost.ImageURL)
	fmt.Printf("path: %s \n", newPost.ImageURL)
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer tmpfile.Close()

	_, err = io.Copy(tmpfile, file)
	if err != nil {
		ctx.Logger.WithError(err).Error("error copying file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}
