package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

var WIDTH = 500
var HEIGHT = 500

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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	// Decodifica l'immagine in base64
	image, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Create the file
	path := "storage/" + fmt.Sprint(profileUserID) + "/tmp_profile_pic.jpeg"
	err = os.WriteFile(path, image, 0644)
	if err != nil {
		ctx.Logger.WithError(err).Error("error saving image")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Crop the image
	newPath, err := saveAndCrop(path, 250, 250)
	if err != nil {
		ctx.Logger.WithError(err).Error("error saving or cropping image")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Delete tmp file
	err = os.Remove(path)
	if err != nil {
		ctx.Logger.WithError(err).Error("error deleting tmp file")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	type ProfilePic struct {
		Path string `json:"path"`
	}

	pic := ProfilePic{Path: newPath}

	// Return the new post
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(pic); err != nil {
		ctx.Logger.WithError(err).Error("error encoding proPic path")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
