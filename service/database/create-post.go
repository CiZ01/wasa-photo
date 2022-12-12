package database

import "fmt"

var query_CREATEPOST = `INSERT INTO Post (postID, userID, postImageURL, caption) VALUES (?, ?, ?, ?)`

func (db *appdbimpl) CreatePost(p Post) (Post, error) {
	// Check last postID in db
	_postID, err := db.GetLastPostID(p.User.UserID)
	if err != nil {
		return p, err
	}
	postID := fmt.Sprint(_postID + 1)
	userID := fmt.Sprint(p.User.UserID)

	p.PostID = postID

	path := "./storage/" + userID + "/posts/" + postID + ".jpg"

	_, err = db.c.Exec(query_CREATEPOST, postID, userID, path, p.Caption)
	if err != nil {
		return p, err
	}

	p.ImageURL = path

	return p, nil
}
