package database

import "database/sql"

var query_DELETEPHOTO = `DELETE FROM Post WHERE ownerID=? AND postID=?`
var query_DELETEPHOTO_LIKES = `DELETE FROM Like WHERE ownerID=? AND postID=?`
var query_DELETEPHOTO_COMMENTS = `DELETE FROM Comment WHERE ownerID=? AND postID=?`

/*
DeletePost deletes a post from the database.
Delete all the likes and comments associated with the post.
If the post does not exist, return an error. (?)
*/
func (db *appdbimpl) DeletePost(ownerID int, postID int) error {

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

	// Delete all the likes
	_, err = tx.Exec(query_DELETEPHOTO_LIKES, ownerID, postID)
	if err != nil {
		return err
	}

	// Delete all the comments
	_, err = tx.Exec(query_DELETEPHOTO_COMMENTS, ownerID, postID)
	if err != nil {
		return err
	}

	// Delete the post
	_, err = tx.Exec(query_DELETEPHOTO, ownerID, postID)
	if err != nil {
		return err
	}
	return nil
}
