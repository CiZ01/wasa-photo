package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/utils"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
UploadPhoto is the handler for the POST /users/:profileUserID/posts endpoint.
It creates a new post
*/
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user is authorized
	if profileUserID != ctx.UserID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Parse the multipart form
	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	caption := r.FormValue("caption")

	// Get the file from the form
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	// Read the file
	data, err := io.ReadAll(file)
	if err != nil {
		ctx.Logger.WithError(err).Error("error parse file")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fileType := http.DetectContentType(data)
	if fileType != "image/jpeg" {
		http.Error(w, "Bad Request wrong file type", http.StatusBadRequest)
		return
	}

	defer func() { err = file.Close() }()

	// Get the user from the database
	dbuser, err := rt.db.GetUserByID(profileUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error getting user")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse the user from the database to the User struct in the api package
	var user User
	err = user.FromDatabase(dbuser)
	if err != nil {
		ctx.Logger.WithError(err).Error("error parsing user")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create the new post
	// The caption is taken from the form, the FormValue function returns an empty string if the key is not present
	var newPost = Post{
		User:    user,
		Caption: caption,
	}

	// Parse the new post from the api package to the Post struct in the database package
	dbPost := newPost.ToDatabase()

	// Create the post in the database
	// Here the new post is returned from the database package
	dbNewPost, err := rt.db.CreatePost(dbPost)
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating post")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Save the image
	err = os.WriteFile(dbNewPost.ImageURL, data, 0666)
	if err != nil {
		ctx.Logger.WithError(err).Error("error saving file")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Crop the image
	err = utils.SaveAndCrop(dbNewPost.ImageURL, 720, 720)
	if err != nil {
		ctx.Logger.WithError(err).Error("error cropping file")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse the new post from the database package to the Post struct in the api package
	err = newPost.FromDatabase(dbNewPost)
	if err != nil {
		ctx.Logger.WithError(err).Error("error parsing photo")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the new post
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newPost); err != nil {
		ctx.Logger.WithError(err).Error("Error while encoding the post")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
