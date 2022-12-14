package database

var query_DELETEFOLLOW = `DELETE FROM Follow WHERE followerID = ? AND followedID = ?`

func (db *appdbimpl) DeleteFollow(followerID uint32, followedID uint32) error {
	_, err := db.c.Exec(query_DELETEFOLLOW, followerID, followedID)
	return err
}
