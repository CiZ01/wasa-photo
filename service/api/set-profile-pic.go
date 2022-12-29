package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/disintegration/imaging"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

var WIDTH = 500
var HEIGHT = 500

func (rt *_router) setMyProfilePic(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the user ID from the URL
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

	reader := bufio.NewReader(file)

	// utilizza la funzione Decode per decodificare l'immagine dal reader
	img, _, err := imaging.Decode(reader)
	if err != nil {
		ctx.Logger.WithError(err).Error("error decoding image")
		http.Error(w, "Interal Server Error", http.StatusInternalServerError)
		return
	}

	img = imaging.Resize(img, WIDTH, HEIGHT, imaging.Lanczos)

	// salva l'immagine downscalata in un nuovo file
	err = imaging.Save(img, "immagine_downscaled.jpg")
	if err != nil {
		// gestisci l'errore
	}

	// Create the file
	path := "storage/" + fmt.Sprint(profileUserID) + "/profile_pic.jpg"
	tmpfile, err := os.Create(path)
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

	type profilePic struct {
		path string `json:"path"`
	}

	pic := profilePic{path: path}

	// Return the new post
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(pic); err != nil {
		ctx.Logger.WithError(err).Error("Error while encoding the post")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

// Da IMPLEMENTARE
