package api

import (
	"encoding/base64"
	"fmt"
	"git.francescofazzari.it/wasa_photo/service/database"
	"io/ioutil"
	"os"
	"time"
)

/*
This struct rappresents the Post object.
The post is identified by the PostID, which is the primary key.
*/
type Post struct {
	PostID       string    `json:"postID"`
	User         User      `json:"user"`
	Image        string    `json:"image"`
	Caption      string    `json:"caption"`
	LikeCount    int       `json:"likeCount"`
	CommentCount int       `json:"commentCount"`
	Liked        bool      `json:"liked"`
	Timestamp    time.Time `json:"timestamp"`
}

/*
This function add the user api struct to the post api struct.
*/
func (p *Post) AddUser(user User) Post {
	return Post{
		PostID:       p.PostID,
		User:         User{UserID: user.UserID, Username: user.Username, UserPropicURL: user.UserPropicURL},
		Image:        p.Image,
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
		ImageURL:     "",
		Caption:      p.Caption,
		LikeCount:    p.LikeCount,
		CommentCount: p.CommentCount,
		Timestamp:    p.Timestamp,
	}
}

func (p *Post) FromDatabase(dbPost database.Post) (Post, error) {
	image64, err := imageToBase64(dbPost.ImageURL)
	if err != nil {
		return Post{}, err
	}
	return Post{
		PostID:       dbPost.PostID,
		User:         User{UserID: dbPost.User.UserID, Username: dbPost.User.Username, UserPropicURL: dbPost.User.UserPropicURL},
		Image:        image64,
		Caption:      dbPost.Caption,
		LikeCount:    dbPost.LikeCount,
		CommentCount: dbPost.CommentCount,
		Timestamp:    dbPost.Timestamp,
	}, nil
}

func imageToBase64(filename string) (string, error) {
	imageFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer imageFile.Close()

	imageData, err := ioutil.ReadAll(imageFile)
	if err != nil {
		return "", err
	}

	base64 := base64.StdEncoding.EncodeToString(imageData)
	fmt.Println(base64)
	return base64, nil
}
