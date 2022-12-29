package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
UploadPhoto is the handler for the POST /users/:profileUserID/posts endpoint.
It creates a new post and returns a JSON representation of the new post and 201 response.
The JSON response body is of the form:
- User: the user who created the post, represented as a JSON object, with the following fields:
  - ID: the user ID
  - Username: the username
  - profilePicURL: the URL of the profile picture

- Caption: the caption of the post
*/
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the profileUserID from the URL
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

	// Parse the multipart form
	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Access the photo key
	// The photo key is the name of the file input in the HTML form
	// If the key is not present an error is returned
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Get the user from the database
	dbuser, err := rt.db.GetUserByID(profileUserID)
	if err != nil {
		ctx.Logger.WithError(err).Error("error getting user")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Parse the user from the database to the User struct in the api package
	var user User
	user.FromDatabase(dbuser)

	// Create the new post
	// The caption is taken from the form, the FormValue function returns an empty string if the key is not present
	var newPost = Post{
		User:    user,
		Caption: string(r.FormValue("caption")),
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

	// Parse the new post from the database package to the Post struct in the api package
	newPost = newPost.FromDatabase(dbNewPost)

	// Create the file
	tmpfile, err := os.Create(newPost.ImageURL)
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating file")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer tmpfile.Close()

	// Copy the uploaded file to the created file
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		ctx.Logger.WithError(err).Error("error copying file")
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
