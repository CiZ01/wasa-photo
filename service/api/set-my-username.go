package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SetMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	paramUserID, _ := strconv.Atoi(ps.ByName("profileUserID"))
	userID := uint32(paramUserID)
	header := strings.Split(r.Header.Get("Authorization"), "")

	if !isAuthorized(userID, header) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	type NewUsernameBody struct {
		NewUsername string `json:"username"`
	}

	var newUsernameBody NewUsernameBody

	err := json.NewDecoder(r.Body).Decode(&newUsernameBody)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	newUsername := newUsernameBody.NewUsername

	if !IsValid(newUsername) {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	err = rt.db.ChangeUsername(userID, newUsername)
	if err != nil {
		http.Error(w, "Username already taken. Username must be unique", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "plain/text")
	json.NewEncoder(w).Encode("Username changed")

}
