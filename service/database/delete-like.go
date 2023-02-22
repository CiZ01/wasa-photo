package database

var query_DELETELIKE = `DELETE FROM Like WHERE userID = ? AND ownerID = ? AND postID = ?`

func (db *appdbimpl) DeleteLike(ownerID int, postID int, userID int) error {
	res, err := db.c.Exec(query_DELETELIKE, userID, ownerID, postID)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return ErrNoRowsAffected
	}

	return nil
}
