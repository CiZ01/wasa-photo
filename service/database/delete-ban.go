package database

import "database/sql"

var query_DELETEBAN = `DELETE FROM Ban WHERE bannerID = ? AND bannedID = ?`
var query_SHOWCOMMENT = `UPDATE Comment SET hidden = FALSE WHERE userID = ? AND postID = ?`

func (db *appdbimpl) DeleteBan(bannerID uint32, bannedID uint32) error {
	// Get all posts from the banner user
	rows, err := db.c.Query(query_GETALLPOST, bannerID)
	if err != nil {
		return err
	}

	var posts []uint32
	for rows.Next() {
		var postID uint32
		err = rows.Scan(&postID)
		if err != nil {
			return err
		}
		posts = append(posts, postID)
	}
	err = rows.Close()
	if err != nil {
		return err
	}

	tx, err := db.c.BeginTx(db.ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Unhide all comments from the banned user in banner's posts
	showComment, err := tx.Prepare(query_SHOWCOMMENT)
	for _, postID := range posts {
		// Unhide the comment
		_, err := showComment.Exec(bannedID, postID)
		if err != nil {
			return err
		}
	}

	// Delete the ban relationship between the banner and the banned users.
	_, err = tx.Exec(query_DELETEBAN, bannerID, bannedID)
	if err != nil {
		return err
	}
	return nil
}
