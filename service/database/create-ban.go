package database

var query_CREATEBAN = `INSERT INTO Ban VALUES (?, ?)`
var query_GETALLPOST = `SELECT postID FROM Post WHERE userID=?`
var query_HIDECOMMENTS = `UPDATE Comment SET hidden = "1" WHERE userID=? AND postID=?`
var query_DELETEALLLIKES = `DELETE FROM Like WHERE userID = ? AND postID = ?`

func (db *appdbimpl) CreateBan(bannerID uint32, bannedID uint32) error {
	rows, err := db.c.Query(query_GETALLPOST, bannerID)
	if err != nil {
		return err
	}

	var allPosts []uint32

	for rows.Next() {
		postID := uint32(0)
		err = rows.Scan(&postID)
		if err != nil {
			return err
		}
		allPosts = append(allPosts, postID)
	}
	rows.Close()

	hideComments, err := db.c.Prepare(query_HIDECOMMENTS)
	if err != nil {
		return err
	}

	deleteLikes, err := db.c.Prepare(query_DELETEALLLIKES)
	if err != nil {
		return err
	}

	for _, postID := range allPosts {
		// Hide comments
		hideComments.Exec(bannedID, postID)
		// Delete likes
		deleteLikes.Exec(bannedID, postID)
	}
	deleteLikes.Close()
	hideComments.Close()

	// Delete follow
	// If the user dont follow the banned user, it will return an error
	err = db.DeleteFollow(bannerID, bannedID)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(query_CREATEBAN, bannerID, bannedID)
	if err != nil {
		return err
	}
	return nil
}
