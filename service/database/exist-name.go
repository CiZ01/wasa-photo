package database

var query_FINDUSERNAME = `SELECT username FROM User
						WHERE username = ?`

func (db *appdbimpl) ExistsName(username string) bool {
	var existsName string
	_ = db.c.QueryRow(query_FINDUSERNAME, username).Scan(&existsName)
	return existsName != ""
}
