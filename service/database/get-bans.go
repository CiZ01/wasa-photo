package database

var query_GETBANS = `SELECT bannedID FROM Ban WHERE bannerID = ? LIMIT ?,?`

func (db *appdbimpl) GetBans(bannerID uint32, offset uint32, limit uint32) ([]User, error) {
	var bans []User
	rows, err := db.c.Query(query_GETBANS, bannerID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
	return bans, nil
}
