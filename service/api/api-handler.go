package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	rt.router.POST("/login", rt.wrap(rt.createUser))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
