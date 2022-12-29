package database

var query_SETPROFILEPIC = `UPDATE User SET userPropicURL=? WHERE userID=?`

func (db *appdbimpl) SetProfilePic(userID uint32, profilePicURL string) error {
	_, err := db.c.Exec(query_SETPROFILEPIC, profilePicURL, userID)
	if err != nil {
		return err
	}
	return nil
}
