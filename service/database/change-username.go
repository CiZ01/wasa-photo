package database

var query_CHANGEUSERNAME = `UPDATE User SET username = ? WHERE userID = ?;`

func (db *appdbimpl) ChangeUsername(userID uint32, newUsername string) error {
	_, err := db.c.Exec(query_CHANGEUSERNAME, newUsername, userID)
	if err != nil {
		return err
	}
	return nil
}
