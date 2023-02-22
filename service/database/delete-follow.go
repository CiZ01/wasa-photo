package database

var query_DELETEFOLLOW = `DELETE FROM Follow WHERE followerID = ? AND followedID = ?`

func (db *appdbimpl) DeleteFollow(followerID int, followedID int) error {
	_, err := db.c.Exec(query_DELETEFOLLOW, followerID, followedID)
	return err
}
