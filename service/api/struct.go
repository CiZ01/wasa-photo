package api

import (
	"time"
)

/*
This struct rappresents the Profile object.
The profile is identified by the User.UserID, which is the primary key.
*/
type Profile struct {
	User            User   `json:"user"`
	UserPropicURL   string `json:"userPropicURL"`
	Bio             string `json:"bio"`
	FollowerCount   int    `json:"followerCount"`
	FollowingsCount int    `json:"followingsCount"`
}

/*
This struct rappresents the Comment object.
The post is identified by the CommentID, which is the primary key.
*/
type Comment struct {
	CommentID string    `json:"commentID"`
	User      User      `json:"user"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
	//non so se sia il massimo usare time visto che
	//il timestamp lo gestirò lato frontend e comunque
	//quando gli arriva è una stringa
}
