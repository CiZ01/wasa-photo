package database

var query_GETBANS = `SELECT bannedID FROM Ban WHERE bannerID = ? LIMIT ?,?`

/*
GetBans returns a list of users that have been banned by the user with the given ID.
*/
func (db *appdbimpl) GetBans(bannerID uint32, offset uint32, limit int32) ([]User, error) {
	var bans []User
	rows, err := db.c.Query(query_GETBANS, bannerID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer func() { err = rows.Close() }()

	for rows.Next() {
		var bannedID uint32
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
