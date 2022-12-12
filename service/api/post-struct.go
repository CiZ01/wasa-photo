package api

import (
	"time"

	"git.francescofazzari.it/wasa_photo/service/database"
)

/*
This struct rappresents the Post object.
The post is identified by the PostID, which is the primary key.
*/
type Post struct {
	PostID       string    `json:"postID"`
	User         User      `json:"user"`
	ImageURL     string    `json:"imageURL"`
	Caption      string    `json:"caption"`
	LikeCount    int       `json:"likeCount"`
	CommentCount int       `json:"commentCount"`
	Timestamp    time.Time `json:"timestamp"`
}

func (p *Post) ToDatabase(user User) database.Post {
	return database.Post{
		PostID:       p.PostID,
		User:         database.User{UserID: user.UserID, Username: user.Username, UserPropicURL: user.UserPropicURL},
		ImageURL:     p.ImageURL,
		Caption:      p.Caption,
		LikeCount:    p.LikeCount,
		CommentCount: p.CommentCount,
		Timestamp:    p.Timestamp,
	}
}

func (p *Post) FromDatabase(dbPost database.Post, dbUser database.User) Post {
	return Post{
		PostID:       dbPost.PostID,
		User:         User{UserID: dbUser.UserID, Username: dbUser.Username, UserPropicURL: dbUser.UserPropicURL},
		ImageURL:     dbPost.ImageURL,
		Caption:      dbPost.Caption,
		LikeCount:    dbPost.LikeCount,
		CommentCount: dbPost.CommentCount,
		Timestamp:    dbPost.Timestamp,
	}
}
