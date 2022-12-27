package database

import (
	"database/sql"
	"fmt"
	"os"
)

var query_DELETEALLLIKE = `DELETE FROM Like WHERE userID = ?`
var query_DELETEALLCOMMENT = `DELETE FROM Comment WHERE userID = ?`
var query_DELETEALLFOLLOW = `DELETE FROM Follow WHERE userID = ? OR targetUserID = ?`
var query_DELETEALLBAN = `DELETE FROM Ban WHERE userID = ? OR targetUserID = ?`
var query_DELETEALLPOSTS = `DELETE FROM Post WHERE userID = ?`
var query_DELETEUSER = `DELETE FROM User WHERE userID = ?`

func (db *appdbimpl) DeleteUser(userID uint32) error {

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

	_, err = tx.Exec(query_DELETEALLPOSTS, userID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(query_DELETEALLBAN, userID, userID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(query_DELETEALLFOLLOW, userID, userID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(query_DELETEALLCOMMENT, userID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(query_DELETEALLLIKE, userID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(query_DELETEUSER, userID)
	if err != nil {
		return err
	}

	err = os.RemoveAll("./storage/" + fmt.Sprint(userID) + "/")
	if err != nil {
		return err
	}

	return err
}
