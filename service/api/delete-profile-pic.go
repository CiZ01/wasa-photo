package api

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"git.francescofazzari.it/wasa_photo/service/api/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) resetMyProfilePic(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the user ID from the URL
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

	dbUser, err := rt.db.GetUserByID(userID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while getting the user")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var user User
	err = user.FromDatabase(dbUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while converting the user")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	source, err := os.Open("./storage/default_propic.jpg")
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while opening the default profile pic")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer source.Close()

	destination, err := os.Create(utils.GetProfilePicPath(user.UserID))
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while creating the profile pic")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error while copying the default profile pic")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
