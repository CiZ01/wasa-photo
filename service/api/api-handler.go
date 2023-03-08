package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// -------LOGIN AND REGISTER--------//
	rt.router.POST("/session", rt.wrap(rt.doLogin, false))

	// ----------DELETE USER-------------//
	rt.router.DELETE("/profiles/:profileUserID", rt.wrap(rt.deleteUser, true))

	// --------GET USER PROFILE----------//
	rt.router.GET("/profiles/:profileUserID", rt.wrap(rt.getUserProfile, true))

	// --------CHANGE USERNAME---------//
	rt.router.PUT("/profiles/:profileUserID/username", rt.wrap(rt.setMyUserName, true))

	// --------CHANGE BIO---------------//
	rt.router.PUT("/profiles/:profileUserID/bio", rt.wrap(rt.setMyBio, true))

	// --------CHANGE PROFILE PIC-------//
	rt.router.PUT("/profiles/:profileUserID/profile-picture", rt.wrap(rt.setMyProfilePic, true))

	// ----------FOLLOW USER-----------//
	rt.router.PUT("/profiles/:profileUserID/followings/:targetUserID", rt.wrap(rt.followUser, true))

	// --------UNFOLLOW USER-----------//
	rt.router.DELETE("/profiles/:profileUserID/followings/:targetUserID", rt.wrap(rt.unfollowUser, true))

	// ----------GET FOLLOWINGS--------//
	rt.router.GET("/profiles/:profileUserID/followings", rt.wrap(rt.getMyFollowings, true))

	// ----------GET FOLLOWERS---------//
	rt.router.GET("/profiles/:profileUserID/followers", rt.wrap(rt.getMyFollowers, true))

	// ----------BAN USER--------------//
	rt.router.PUT("/profiles/:profileUserID/bans/:targetUserID", rt.wrap(rt.banUser, true))

	// --------UNBAN USER--------------//
	rt.router.DELETE("/profiles/:profileUserID/bans/:targetUserID", rt.wrap(rt.unbanUser, true))

	// -----------GET BANS------------//
	rt.router.GET("/profiles/:profileUserID/bans", rt.wrap(rt.getMyBans, true))

	// ----------UPLOAD POST----------//
	rt.router.POST("/profiles/:profileUserID/posts", rt.wrap(rt.uploadPhoto, true))

	// ----------GET POSTS------------//
	rt.router.GET("/profiles/:profileUserID/posts", rt.wrap(rt.getPosts, true))

	// ----------CHANGE CAPTION-------------//
	rt.router.PUT("/profiles/:profileUserID/posts/:postID/caption", rt.wrap(rt.changeCaption, true))

	// ----------DELETE POST-----------//
	rt.router.DELETE("/profiles/:profileUserID/posts/:postID", rt.wrap(rt.deletePhoto, true))

	// ----------LIKE POST-------------//
	rt.router.PUT("/profiles/:profileUserID/posts/:postID/likes/:userID", rt.wrap(rt.likePhoto, true))

	// --------UNLIKE POST-------------//
	rt.router.DELETE("/profiles/:profileUserID/posts/:postID/likes/:userID", rt.wrap(rt.unlikePhoto, true))

	// ---------GET LIKES--------------//
	rt.router.GET("/profiles/:profileUserID/posts/:postID/likes", rt.wrap(rt.getLikes, true))

	// ----------COMMENT POST----------//
	rt.router.POST("/profiles/:profileUserID/posts/:postID/comments", rt.wrap(rt.commentPhoto, true))

	// ----------UNCOMMENT POST--------//
	rt.router.DELETE("/profiles/:profileUserID/posts/:postID/comments/:commentID", rt.wrap(rt.uncommentPhoto, true))

	// ---------GET COMMENTS-----------//
	rt.router.GET("/profiles/:profileUserID/posts/:postID/comments", rt.wrap(rt.getComments, true))

	// ---------GET FEED--------------//
	rt.router.GET("/profiles/:profileUserID/feed", rt.wrap(rt.getMyStream, true))

	// -----------SEARCH------------//
	rt.router.GET("/profiles", rt.wrap(rt.searchUsers, true))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
