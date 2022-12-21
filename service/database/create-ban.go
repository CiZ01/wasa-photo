package database

import (
	"database/sql"
)

var query_CREATEBAN = `INSERT INTO Ban VALUES (?, ?)`
var query_GETALLPOST = `SELECT postID FROM Post WHERE userID=?`
var query_HIDECOMMENTS = `UPDATE Comment SET hidden = "1" WHERE userID=? AND postID=?`
var query_DELETEALLLIKES = `DELETE FROM Like WHERE userID = ? AND postID = ?`

func (db *appdbimpl) CreateBan(bannerID uint32, bannedID uint32) error {
	rows, err := db.c.Query(query_GETALLPOST, bannerID)
	if err != nil {
		return err
	}

	var allPosts []uint32

	for rows.Next() {
		postID := uint32(0)
		err = rows.Scan(&postID)
		if err != nil {
			return err
		}
		allPosts = append(allPosts, postID)
	}
	rows.Close()

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

	hideComments, err := tx.Prepare(query_HIDECOMMENTS)
	if err != nil {
		return err
	}

	deleteLikes, err := tx.Prepare(query_DELETEALLLIKES)
	if err != nil {
		return err
	}

	for _, postID := range allPosts {
		// Hide comments
		_, err := hideComments.Exec(bannedID, postID)
		if err != nil {
			return err
		}
		// Delete likes
		_, err = deleteLikes.Exec(bannedID, postID)
		if err != nil {
			return err
		}
	}
	deleteLikes.Close()
	hideComments.Close()

	// Delete follow
	// If the user dont follow the banned user, it will return an error
	err = db.DeleteFollow(bannerID, bannedID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(query_CREATEBAN, bannerID, bannedID)
	if err != nil {
		return err
	}
	return nil
}
