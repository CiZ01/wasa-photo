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
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&isBanned)
	}

	return isBanned != "", nil
}
