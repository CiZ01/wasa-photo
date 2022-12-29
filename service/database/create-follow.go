package database

var query_CREATEFOLLOW = `INSERT INTO Follow (followerID, followedID) VALUES (?, ?)`

func (db *appdbimpl) CreateFollow(followerID uint32, followedID uint32) error {
	_, err := db.c.Exec(query_CREATEFOLLOW, followerID, followedID)
	return err
}
