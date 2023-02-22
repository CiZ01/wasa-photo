package database

var query_SETPROFILEPIC = `UPDATE User SET userPropicURL=? WHERE userID=?`

func (db *appdbimpl) SetProfilePic(userID int, profilePicURL string) error {
	_, err := db.c.Exec(query_SETPROFILEPIC, profilePicURL, userID)
	return err
}
