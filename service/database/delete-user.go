package database

import (
	"database/sql"
	"fmt"
	"os"
)

var query_DELETEUSER = `DELETE FROM User WHERE userID = ?`

func (db *appdbimpl) DeleteUser(userID int) error {

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

	// Delete the user
	_, err = tx.Exec(query_DELETEUSER, userID)
	if err != nil {
		return err
	}

	// Delete files
	err = os.RemoveAll("./storage/" + fmt.Sprint(userID) + "/")
	if err != nil {
		return err
	}

	return err
}
