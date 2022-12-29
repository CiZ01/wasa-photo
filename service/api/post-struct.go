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

/*
This function add the user api struct to the post api struct.
*/
func (p *Post) AddUser(user User) Post {
	return Post{
		PostID:       p.PostID,
		User:         User{UserID: user.UserID, Username: user.Username, UserPropicURL: user.UserPropicURL},
		ImageURL:     p.ImageURL,
		Caption:      p.Caption,
		LikeCount:    p.LikeCount,
		CommentCount: p.CommentCount,
		Timestamp:    p.Timestamp,
	}
}

/*
This function parse the post api struct to the post database struct.
Also the user api struct is parsed to the user database struct.
*/
func (p *Post) ToDatabase() database.Post {
	return database.Post{
		PostID:       p.PostID,
		User:         database.User{UserID: p.User.UserID, Username: p.User.Username, UserPropicURL: p.User.UserPropicURL},
		ImageURL:     p.ImageURL,
		Caption:      p.Caption,
		LikeCount:    p.LikeCount,
		CommentCount: p.CommentCount,
		Timestamp:    p.Timestamp,
	}
}

func (p *Post) FromDatabase(dbPost database.Post) Post {
	return Post{
		PostID:       dbPost.PostID,
		User:         User{UserID: dbPost.User.UserID, Username: dbPost.User.Username, UserPropicURL: dbPost.User.UserPropicURL},
		ImageURL:     dbPost.ImageURL,
		Caption:      dbPost.Caption,
		LikeCount:    dbPost.LikeCount,
		CommentCount: dbPost.CommentCount,
		Timestamp:    dbPost.Timestamp,
	}
}
