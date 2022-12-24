package database

var query_GETBAN = `SELECT bannerID FROM Ban WHERE 
(bannerID = ? AND bannedID = ?) OR
(bannerID = ? AND bannedID = ?)`

func (db *appdbimpl) IsBanned(bannedID uint32, bannerID uint32) (bool, error) {
	var isBanned string
	rows, err := db.c.Query(query_GETBAN, bannerID, bannedID, bannedID, bannerID)
	if err != nil {
		return false, err
	}
	defer func() { err = rows.Close() }()

	for rows.Next() {
		err := rows.Scan(&isBanned)
		if err != nil {
			return false, err
		}
	}

	return isBanned != "", err
}
