package database

import (
	"database/sql"
	"os"

	"git.francescofazzari.it/wasa_photo/service/api/utils"
)

var query_DELETEPHOTO = `DELETE FROM Post WHERE userID=? AND postID=?`

/*
DeletePost deletes a post from the database.
Delete all the likes and comments associated with the post.
If the post does not exist, return an error. (?)
*/
func (db *appdbimpl) DeletePost(userID int, postID int) error {

	tx, err := db.c.BeginTx(db.ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
		}
		err = tx.Commit()
	}()

	// Delete the post
	_, err = tx.Exec(query_DELETEPHOTO, userID, postID)
	if err != nil {
		return err
	}

	// Delete file
	err = os.Remove(utils.GetPostPhotoPath(userID, postID))
	if err != nil {
		return err
	}

	return err
}
