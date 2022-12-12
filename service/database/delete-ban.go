package database

var sql_DELETEBAN = `DELETE FROM Ban WHERE bannerID = ? AND bannedID = ?`

func (db *appdbimpl) DeleteBan(bannerID uint32, bannedID uint32) error {
	statement, err := db.c.Prepare(sql_DELETEBAN)
	if err != nil {
		return err
	}
	statement.Exec(bannerID, bannedID)
	statement.Close()
	return nil
}
