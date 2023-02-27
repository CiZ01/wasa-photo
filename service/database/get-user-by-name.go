package database

var query_GETUSERBYNAME = `SELECT userID, username FROM User WHERE username = ?;`

/*
GetUSer function returns the user object from the database.
*/
func (db *appdbimpl) GetUserByName(username string) (User, error) {
	var user User
	err := db.c.QueryRow(query_GETUSERBYNAME, username).Scan(&user.UserID, &user.Username)
	return user, err
}
