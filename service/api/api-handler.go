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
	rt.router.POST("/profiles/:profileUserID/posts", rt.wrap(rt.UploadPhoto))

	// ----------GET POSTS------------//
	rt.router.GET("/profiles/:profileUserID/posts", rt.wrap(rt.GetPosts))

	// ----------LIKE POST-------------//
	rt.router.PUT("/profiles/:profileUserID/posts/:postID/likes/:userID", rt.wrap(rt.LikePhoto))

	// --------UNLIKE POST-------------//
	rt.router.DELETE("/profiles/:profileUserID/posts/:postID/likes/:userID", rt.wrap(rt.UnlikePhoto))

	// ---------GET LIKES--------------//
	rt.router.GET("/profiles/:profileUserID/posts/:postID/likes", rt.wrap(rt.GetLikes))

	// ----------COMMENT POST----------//
	rt.router.POST("/profiles/:profileUserID/posts/:postID/comments/:userID", rt.wrap(rt.CommentPhoto))

	// ----------UNCOMMENT POST--------//
	rt.router.DELETE("/profiles/:profileUserID/posts/:postID/comments/:commentID", rt.wrap(rt.UncommentPhoto))

	// ---------GET COMMENTS-----------//
	rt.router.GET("/profiles/:profileUserID/posts/:postID/comments", rt.wrap(rt.GetComments))

	// ----------DELETE USER-------------//
	rt.router.DELETE("/profiles/:profileUserID", rt.wrap(rt.DeleteUser))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
