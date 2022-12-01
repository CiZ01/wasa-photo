package database

var query_ADDUSER = `INSERT INTO User (userID, username, userPropicURL)
					 VALUES (?, ?, ?)`

func (db *appdbimpl) CreateUser(u User) (User, error) {
	res, err := db.c.Exec(query_ADDUSER,
		u.UserID, u.Username, u.UserPropicURL)
	if err != nil {
		return u, err
	}

	//questo mi serve
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}

	u.UserID = uint32(lastInsertID)
	return u, nil
}
