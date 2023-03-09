package api

import (
	"encoding/json"
	"net/http"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
doLogin is the handler for the API endpoint POST /session.
It takes the username from the request body and returns the user object and the authoization token in a JSON object.
If the user does not exist, it creates a new user.
The request body must be a JSON object with the following fields:
- username: string
*/
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User

	// Read the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the username is valid
	if !IsValid(user.Username) {
		http.Error(w, "Bad reequest invalid username", http.StatusBadRequest)
		return
	}

	/*
		Check if the user exists
		If the user does not exist, create a new user
		else return the user object
	*/
	exist, err := rt.db.ExistsName(user.Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't check if the user exists")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !exist {
		user, err = rt.CreateUser(user)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't create the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		dbUser, err := rt.db.GetUserByName(user.Username)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't load the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = user.FromDatabase(dbUser)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't convert the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	// This struct contain the User object and the authorization token.
	type AuthUser struct {
		User  User `json:"user"`
		Token int  `json:"token"`
	}

	/*
		Create the AuthUser object and set the user object and the authorization token.
		The authorization token is the userID.
	*/
	authUser := AuthUser{user, user.UserID}

	// Encode the AuthUser object in JSON and send it to the client.
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(authUser); err != nil {
		ctx.Logger.WithError(err).Error("can't encode the response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
