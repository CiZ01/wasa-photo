package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// -------LOGIN AND REGISTER--------//
	rt.router.POST("/session", rt.wrap(rt.doLogin))

	// ----------DELETE USER-------------//
	rt.router.DELETE("/profiles/:profileUserID", rt.wrap(rt.deleteUser))

	// --------GET USER PROFILE----------//
	rt.router.GET("/profiles/:profileUserID", rt.wrap(rt.getMyUserProfile))

	// --------CHANGE USERNAME---------//
	rt.router.PUT("/profiles/:profileUserID/username", rt.wrap(rt.setMyUsername))

	// --------CHANGE BIO---------------//
	rt.router.PUT("/profiles/:profileUserID/bio", rt.wrap(rt.setMyBio))

	// --------CHANGE PROFILE PIC-------//
	//rt.router.PUT("/profiles/:profileUserID/profilePic", rt.wrap(rt.setMyProfilePic))

	// ----------FOLLOW USER-----------//
	rt.router.PUT("/profiles/:profileUserID/followings/:targetUserID", rt.wrap(rt.followUser))
	// --------------------------------//

	// --------UNFOLLOW USER-----------//
	rt.router.DELETE("/profiles/:profileUserID/followings/:targetUserID", rt.wrap(rt.unfollowUser))

	// ----------GET FOLLOWINGS--------//
	rt.router.GET("/profiles/:profileUserID/followings", rt.wrap(rt.getMyFollowings))

	// ----------GET FOLLOWERS---------//
	rt.router.GET("/profiles/:profileUserID/followers", rt.wrap(rt.getMyFollowers))

	// ----------BAN USER--------------//
	rt.router.PUT("/profiles/:profileUserID/bans/:targetUserID", rt.wrap(rt.banUser))

	// --------UNBAN USER--------------//
	rt.router.DELETE("/profiles/:profileUserID/bans/:targetUserID", rt.wrap(rt.unbanUser))

	// -----------GET BANS------------//
	rt.router.GET("/profiles/:profileUserID/bans", rt.wrap(rt.getMyBans))

	// ----------UPLOAD POST----------//
	rt.router.POST("/profiles/:profileUserID/posts", rt.wrap(rt.uploadPhoto))

	// ----------GET POSTS------------//
	rt.router.GET("/profiles/:profileUserID/posts", rt.wrap(rt.getPosts))

	// ----------DELETE POST-----------//
	rt.router.DELETE("/profiles/:profileUserID/posts/:postID", rt.wrap(rt.deletePhoto))

	// ----------LIKE POST-------------//
	rt.router.PUT("/profiles/:profileUserID/posts/:postID/likes/:userID", rt.wrap(rt.likePhoto))

	// --------UNLIKE POST-------------//
	rt.router.DELETE("/profiles/:profileUserID/posts/:postID/likes/:userID", rt.wrap(rt.unlikePhoto))

	// ---------GET LIKES--------------//
	rt.router.GET("/profiles/:profileUserID/posts/:postID/likes", rt.wrap(rt.getLikes))

	// ----------COMMENT POST----------//
	rt.router.POST("/profiles/:profileUserID/posts/:postID/comments/:userID", rt.wrap(rt.commentPhoto))

	// ----------UNCOMMENT POST--------//
	rt.router.DELETE("/profiles/:profileUserID/posts/:postID/comments/:commentID", rt.wrap(rt.uncommentPhoto))

	// ---------GET COMMENTS-----------//
	rt.router.GET("/profiles/:profileUserID/posts/:postID/comments", rt.wrap(rt.getComments))

	// ---------GET FEED--------------//
	// rt.router.GET("/profiles/:profileUserID/feed", rt.wrap(rt.getMyStream))

	// -----------GET USERS------------//
	// the search function is still under development, however, it will be called `searchUsers`

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
