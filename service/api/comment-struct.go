package api

import (
	"regexp"
	"time"

	"git.francescofazzari.it/wasa_photo/service/database"
)

/*
This struct rappresents the Comment object.
The post is identified by the CommentID, which is the primary key.
*/
type Comment struct {
	CommentID int       `json:"commentID"`
	PostID    int       `json:"postID"`
	OwnerID   int       `json:"ownerID"`
	User      User      `json:"user"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func (c *Comment) FromDatabase(dbcomment database.Comment) error {
	c.CommentID = dbcomment.CommentID
	c.Text = dbcomment.Text
	c.Timestamp = dbcomment.Timestamp
	c.PostID = dbcomment.PostID
	c.OwnerID = dbcomment.OwnerID

	var user User
	err := user.FromDatabase(dbcomment.User)
	if err != nil {
		return err
	}
	c.User = user
	return nil
}

func (c *Comment) isValid() bool {
	text := c.Text
	validComment := regexp.MustCompile(`^[^\/\\]{1,64}$`)
	return validComment.MatchString(text)

}
