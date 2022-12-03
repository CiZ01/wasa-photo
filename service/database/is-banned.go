package database

var query_GETBAN = `SELECT * FROM Ban WHERE bannerID = ? AND bannedID = ?`

func (db *appdbimpl) IsBanned(bannedID uint32, bannerID uint32) (bool, error) {
	res, err := db.c.Query(query_GETBAN, bannerID, bannedID)
	if err != nil {
		return false, err
	}

	var isBanned string
	res.Scan(&isBanned)
	return isBanned != "", nil
}
