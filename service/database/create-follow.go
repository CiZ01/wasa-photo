package database

var query_CREATEFOLLOW = `INSERT INTO Follow VALUES (?, ?)`

func (db *appdbimpl) CreateFollow(followerID uint32, followedID uint32) error {
	statement, err := db.c.Prepare(query_CREATEFOLLOW)
	statement.Exec(followerID, followedID)
	return err
}
