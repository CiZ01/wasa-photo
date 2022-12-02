package database

var query_GETUSERBYID = `SELECT userID, username, userPropicURL FROM User WHERE userID = ?;`

/*
GetUserByID returns a user with the given userID
If no user is found, an user object with userID = 0 is returned
*/
func (db *appdbimpl) GetUserByID(userID uint32) (User, error) {
	var user User

	err := db.c.QueryRow(query_GETUSERBYID, userID).Scan(&user.UserID, &user.Username, &user.UserPropicURL)
	if err != nil {
		return user, err
	}
	return user, nil
}
