package database

var query_GETUSER = `SELECT userID, username, userPropicURL FROM User WHERE username = ?;`

/*
GetUSer function returns the user object from the database.
*/
func (db *appdbimpl) GetUser(username string) (User, error) {
	var user User

	err := db.c.QueryRow(query_GETUSER, username).Scan(&user.UserID, &user.Username, &user.UserPropicURL)
	if err != nil {
		return user, err
	}
	return user, nil
}
