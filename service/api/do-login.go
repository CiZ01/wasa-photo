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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the username is valid
	if !IsValid(user.Username) {
		w.Header().Set("content-type", "plain/text")
		http.Error(w, "invalid username", http.StatusBadRequest)
		return
	}

	/*
		Check if the user exists
		If the user does not exist, create a new user
		else return the user object
	*/
	if !rt.db.ExistsName(user.Username) {
		user, err = rt.CreateUser(user)
		w.WriteHeader(201)
	} else {
		dbUser, _ := rt.db.GetUserByName(user.Username)
		user.FromDatabase(dbUser)
		w.WriteHeader(200)
	}
	if err != nil {
		ctx.Logger.WithError(err).Error("can't load or create the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// This struct contain the User object and the authorization token.
	type AuthUser struct {
		User User
		Auth uint32
	}

	/*
		Create the AuthUser object and set the user object and the authorization token.
		The authorization token is the userID.
	*/
	authUser := AuthUser{user, user.UserID}

	// Encode the AuthUser object in JSON and send it to the client.
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(authUser)
}
