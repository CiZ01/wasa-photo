package api

import (
	"regexp"
	"time"

	"git.francescofazzari.it/wasa_photo/service/api/utils"
	"git.francescofazzari.it/wasa_photo/service/database"
)

/*
This struct rappresents the Post object.
The post is identified by the PostID, which is the primary key.
*/
type Post struct {
	PostID        int       `json:"postID"`
	User          User      `json:"user"`
	Image         string    `json:"image"`
	Caption       string    `json:"caption"`
	LikesCount    int       `json:"likesCount"`
	CommentsCount int       `json:"commentsCount"`
	Liked         bool      `json:"liked"`
	Timestamp     time.Time `json:"timestamp"`
}

/*
This function add the user api struct to the post api struct.
*/
func (p *Post) AddUser(user User) Post {
	return Post{
		PostID:        p.PostID,
		User:          User{UserID: user.UserID, Username: user.Username, UserPropic64: user.UserPropic64},
		Image:         p.Image,
		Caption:       p.Caption,
		LikesCount:    p.LikesCount,
		CommentsCount: p.CommentsCount,
		Timestamp:     p.Timestamp,
	}
}

/*
This function parse the post api struct to the post database struct.
Also the user api struct is parsed to the user database struct.
*/
func (p *Post) ToDatabase() database.Post {
	return database.Post{
		PostID:        p.PostID,
		User:          database.User{UserID: p.User.UserID, Username: p.User.Username},
		Caption:       p.Caption,
		LikesCount:    p.LikesCount,
		CommentsCount: p.CommentsCount,
		Liked:         p.Liked,
		Timestamp:     p.Timestamp,
	}
}

func (p *Post) FromDatabase(dbPost database.Post) error {
	image64, err := utils.ImageToBase64(utils.GetPostPhotoPath(dbPost.User.UserID, dbPost.PostID))
	if err != nil {
		return err
	}

	var apiUser User
	err = apiUser.FromDatabase(dbPost.User)
	if err != nil {
		return err
	}

	p.PostID = dbPost.PostID
	p.User = apiUser
	p.Image = image64
	p.Caption = dbPost.Caption
	p.LikesCount = dbPost.LikesCount
	p.CommentsCount = dbPost.CommentsCount
	p.Liked = dbPost.Liked
	p.Timestamp = dbPost.Timestamp
	return nil
}

func (p *Post) isValid() bool {
	caption := p.Caption
	validCaption := regexp.MustCompile(`^[^\/\\]{0,64}$`)
	return validCaption.MatchString(caption)
}
