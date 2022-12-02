package api

import (
	"encoding/json"
	"net/http"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
createUser read the request body and create a new account.
The request body must be a JSON object with the following fields:
- username: string
*/
func (rt *_router) createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User

	//teoricamente così carica anche json fatti male, basta che ci sia username
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !user.IsValid() {
		//manca l'header
		http.Error(w, "invalid username", http.StatusBadRequest)
		return
	}

	dbUser, err := rt.db.CreateUser(user.ToDatabase())
	if err != nil {
		//forse qua devo cambiare la head perché su api ho messo che ritorno plain/text
		ctx.Logger.WithError(err).Error("can't create the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//carico lo user con l'id che mi ha dato il db
	var currentUser User
	currentUser.FromDatabase(dbUser)

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(currentUser)
}
