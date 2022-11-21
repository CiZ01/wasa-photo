package api

import (
	"encoding/json"
	"net/http"
	"os"

	"git.francescofazzari.it/wasa_photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// getContextReply is an example of HTTP endpoint that returns "Hello World!" as a plain text. The signature of this
// handler accepts a reqcontext.RequestContext (see httpRouterHandler).
func (rt *_router) getContextReply(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")
	json.NewEncoder(os.Stdout).Encode(ctx)
	_, _ = w.Write([]byte("Hello World!"))
}
