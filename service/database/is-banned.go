package database

var query_GETBAN = `SELECT bannerID FROM Ban WHERE 
(bannerID = ? AND bannedID = ?) OR
(bannedID = ? AND bannerID = ?)`

func (db *appdbimpl) IsBanned(bannedID int, bannerID int) (bool, error) {
	var isBanned string
	rows, err := db.c.Query(query_GETBAN, bannerID, bannedID, bannerID, bannedID)
	if err != nil {
		return false, err
	}
	defer func() { err = rows.Close() }()

	for rows.Next() {
		if rows.Err() != nil {
			return false, err
		}
		err := rows.Scan(&isBanned)
		if err != nil {
			return false, err
		}
	}

	return isBanned != "", err
}
