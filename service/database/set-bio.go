package database

var query_SETBIO = `UPDATE User SET bio=? WHERE userID=?`

func (db *appdbimpl) SetBio(userID uint32, bio string) error {
	_, err := db.c.Exec(query_SETBIO, bio, userID)
	if err != nil {
		return err
	}
	return nil
}
