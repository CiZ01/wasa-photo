package database

var query_CHANGEUSERNAME = `UPDATE User SET username = ? WHERE userID = ?;`

func (db *appdbimpl) ChangeUsername(userID uint32, newUsername string) error {
	statement, err := db.c.Prepare(query_CHANGEUSERNAME)
	statement.Exec(newUsername, userID)
	statement.Close()
	return err
}
