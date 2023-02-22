package database

var query_CREATELIKE = `INSERT INTO Like (userID, ownerID, postID) VALUES (?, ?, ?)`

func (db *appdbimpl) CreateLike(userID int, ownerID int, postID int) error {
	_, err := db.c.Exec(query_CREATELIKE, userID, ownerID, postID)
	if err != nil {
		return err
	}

	return nil
}
