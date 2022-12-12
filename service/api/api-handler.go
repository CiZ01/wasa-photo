package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// -------LOGIN AND REGISTER--------//
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	// --------------------------------//

	// --------CHANGE USERNAME---------//
	rt.router.PUT("/profiles/:profileUserID/username", rt.wrap(rt.SetMyUsername))
	// --------------------------------//

	// ----------FOLLOW USER-----------//
	rt.router.PUT("/profiles/:profileUserID/followings/:targetUserID", rt.wrap(rt.FollowUser))
	// --------------------------------//

	// --------UNFOLLOW USER-----------//
	rt.router.DELETE("/profiles/:profileUserID/followings/:targetUserID", rt.wrap(rt.UnfollowUser))

	// ----------BAN USER--------------//
	rt.router.PUT("/profiles/:profileUserID/bans/:targetUserID", rt.wrap(rt.BanUser))

	// --------UNBAN USER--------------//
	rt.router.DELETE("/profiles/:profileUserID/bans/:targetUserID", rt.wrap(rt.UnbanUser))

	// ----------UPLOAD POST----------//
	rt.router.POST("/profiles/:profileUserID/posts", rt.wrap(rt.UploadPost))

	// ----------GET POSTS------------//
	rt.router.GET("/profiles/:profileUserID/posts", rt.wrap(rt.GetPosts))

	// ----------DELETE USER-------------//
	rt.router.DELETE("/profiles/:userID", rt.wrap(rt.DeleteUser))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
