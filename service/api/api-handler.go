package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//--------LOGIN AND REGISTER--------//
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	//---------------------------------//

	//---------CHANGE USERNAME---------//
	rt.router.PUT("/profiles/:profileUserID/username", rt.wrap(rt.SetMyUsername))
	//---------------------------------//

	//-----------FOLLOW USER-----------//
	rt.router.PUT("/profiles/:profileUserID/followings/:targetUserID", rt.wrap(rt.FollowUser))
	//---------------------------------//

	//---------UNFOLLOW USER-----------//
	rt.router.DELETE("/profiles/:profileUserID/followings/:targetUserID", rt.wrap(rt.UnfollowUser))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
