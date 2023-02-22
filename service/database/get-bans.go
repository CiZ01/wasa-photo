package database

var query_GETBANS = `SELECT bannedID FROM Ban WHERE bannerID = ? LIMIT ?,?`

/*
GetBans returns a list of users that have been banned by the user with the given ID.
*/
func (db *appdbimpl) GetBans(bannerID int, offset int, limit int) ([]User, error) {
	var bans []User
	rows, err := db.c.Query(query_GETBANS, bannerID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer func() { err = rows.Close() }()

	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}
		var bannedID int
		err = rows.Scan(&bannedID)
		if err != nil {
			return nil, err
		}

		user, err := db.GetUserByID(bannedID)
		if err != nil {
			return nil, err
		}
		bans = append(bans, user)
	}

	return bans, err
}
