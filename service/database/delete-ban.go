package database

import "database/sql"

var query_DELETEBAN = `DELETE FROM Ban WHERE bannerID = ? AND bannedID = ?`
var query_SHOWCOMMENT = `UPDATE Comment SET hidden = FALSE WHERE userID = ? AND postID = ?`

func (db *appdbimpl) DeleteBan(bannerID int, bannedID int) error {
	// Get all posts from the banner user
	rows, err := db.c.Query(query_GETALLPOST, bannerID)
	if err != nil {
		return err
	}

	var posts []int
	for rows.Next() {
		if rows.Err() != nil {
			return err
		}
		var postID int
		err = rows.Scan(&postID)
		if err != nil {
			return err
		}
		posts = append(posts, postID)
	}
	defer func() { err = rows.Close() }()

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
	return err
}
