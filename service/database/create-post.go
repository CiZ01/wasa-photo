package database

import (
	"database/sql"
	"os"

	"git.francescofazzari.it/wasa_photo/service/api/utils"
)

var query_CREATEPOST = `INSERT INTO Post (postID, userID, caption) VALUES (?, ?, ?)`

func (db *appdbimpl) CreatePost(p Post, data []byte) (Post, error) {
	// Check last postID in db
	_postID, err := db.GetLastPostID(p.User.UserID)
	if err != nil {
		return p, err
	}

	tx, err := db.c.BeginTx(db.ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return p, err
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
		}
		err = tx.Commit()
	}()

	p.PostID = _postID + 1
	profileUserID := p.User.UserID
	path := utils.GetPostPhotoPath(profileUserID, p.PostID)

	// Save the image
	err = os.WriteFile(path, data, 0666)
	if err != nil {
		return p, err
	}

	// Crop the image
	err = utils.SaveAndCrop(path, 720, 720)
	if err != nil {
		return p, err
	}

	_, err = db.c.Exec(query_CREATEPOST, p.PostID, p.User.UserID, p.Caption)
	if err != nil {
		return p, err
	}

	newPost, err := db.GetPosts(p.User.UserID, p.User.UserID, 0, 1)
	if err != nil {
		return p, err
	}

	return newPost[0], err
}
