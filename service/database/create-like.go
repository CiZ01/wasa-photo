package database

var query_CREATELIKE = `INSERT INTO Like (userID, ownerID, postID) VALUES (?, ?, ?)`

func (db *appdbimpl) CreateLike(userID uint32, ownerID uint32, postID uint32) error {
	_, err := db.c.Exec(query_CREATELIKE, userID, ownerID, postID)
	if err != nil {
		return err
	}

	return nil
}
