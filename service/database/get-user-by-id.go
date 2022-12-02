package database

var query_GETUSERBYID = `SELECT userID, username, userPropicURL FROM User WHERE userID = ?;`

func (db *appdbimpl) GetUserByID(userID uint32) (User, error) {
	var user User

	err := db.c.QueryRow(query_GETUSERBYID, userID).Scan(&user.UserID, &user.Username, &user.UserPropicURL)
	if err != nil {
		return user, err
	}
	return user, nil

}
