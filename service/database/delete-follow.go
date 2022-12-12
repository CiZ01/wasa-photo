package database

var query_DELETEFOLLOW = `DELETE FROM Follow WHERE followerID = ? AND followedID = ?`

func (db *appdbimpl) DeleteFollow(followerID uint32, followedID uint32) error {
	statement, err := db.c.Prepare(query_DELETEFOLLOW)
	statement.Exec(followerID, followedID)
	statement.Close()
	return err
}
