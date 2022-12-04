package database

var query_DELETEFOLLOW = `DELETE FROM Follow WHERE followerID = ? AND followedID = ?`

func (db *appdbimpl) DeleteFollow(profileUserID uint32, targetUserID uint32) error {
	_, err := db.c.Exec(query_DELETEFOLLOW, profileUserID, targetUserID)
	return err
}
