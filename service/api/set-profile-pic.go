package api

import (
	"encoding/json"
	"git.francescofazzari.it/wasa_photo/service/api/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyProfilePic(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the user ID from the URL
	profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	userID := ctx.UserID
	if profileUserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Parse the multipart form
	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	// Access the photo key
	// The photo key is the name of the file input in the HTML form
	// If the key is not present an error is returned
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		return
	}

	// Read the file
	data, err := ioutil.ReadAll(file)
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


	// Create the file
	path := utils.GetProfilePicPath(userID)
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		ctx.Logger.WithError(err).Error("error saving image")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Crop the image
	err = utils.SaveAndCrop(path, 250, 250)
	if err != nil {
		ctx.Logger.WithError(err).Error("error saving or cropping image")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	type ProfilePic struct {
		ProfilePic64 string `json:"profilePic64"`
	}

	propic64, err := utils.ImageToBase64(path)
	pic := ProfilePic{ProfilePic64: propic64}

	// Return the new post
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(pic); err != nil {
		ctx.Logger.WithError(err).Error("error encoding proPic path")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
