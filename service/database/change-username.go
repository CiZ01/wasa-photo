package database

var query_CHANGEUSERNAME = `UPDATE User SET username = ? WHERE userID = ?;`

func (db *appdbimpl) ChangeUsername(userID int, newUsername string) error {
	_, err := db.c.Exec(query_CHANGEUSERNAME, newUsername, userID)
	return err
}
