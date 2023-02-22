package database

var query_UPDATECAPTION = `UPDATE Post SET caption = ? WHERE postID = ? AND userID = ?`

func (db *appdbimpl) UpdateCaption(userID int, postID int, newCaption string) error {
	_, err := db.c.Exec(query_UPDATECAPTION, newCaption, postID, userID)
	return err
}
