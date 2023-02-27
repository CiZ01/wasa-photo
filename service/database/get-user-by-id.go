package database

var query_GETUSERBYID = `SELECT userID, username FROM User WHERE userID = ?;`

/*
GetUserByID returns a user with the given userID
If no user is found, it returns an error ErrRowNotFound
*/
func (db *appdbimpl) GetUserByID(userID int) (User, error) {
	var user User
	err := db.c.QueryRow(query_GETUSERBYID, userID).Scan(&user.UserID, &user.Username)
	return user, err
}
