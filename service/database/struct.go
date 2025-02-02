package database

import "time"

type User struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
}

/*
This struct rappresents the Post object.
The post is identified by the PostID, which is the primary key.
*/
type Post struct {
	PostID        int       `json:"postID"`
	User          User      `json:"user"`
	ImageURL      string    `json:"imageURL"`
	Caption       string    `json:"caption"`
	LikesCount    int       `json:"likesCount"`
	CommentsCount int       `json:"commentsCount"`
	Liked         bool      `json:"liked"`
	Timestamp     time.Time `json:"timestamp"`
}

/*
This struct rappresents the Comment object.
The post is identified by the CommentID, which is the primary key.
*/
type Comment struct {
	CommentID int       `json:"commentID"`
	PostID    int       `json:"postID"`
	OwnerID   int       `json:"ownerID"`
	User      User      `json:"userID"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

/*
This struct rappresents the Profile object.
The profile is identified by the User.UserID, which is the primary key.
*/
type Profile struct {
	User            User   `json:"user"`
	Bio             string `json:"bio"`
	FollowersCount  int    `json:"followersCount"`
	FollowingsCount int    `json:"followingsCount"`
	PostsCount      int    `json:"postsCount"`
	IsFollowed      int    `json:"isFollowing"`
}
