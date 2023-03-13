package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) checkForeignKeys(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rt.db.CheckForeignKeys()
}
