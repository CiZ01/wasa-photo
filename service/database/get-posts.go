package database

var query_GETPOSTS = `SELECT postID, userID, postImageURL, caption, timestamp FROM Post WHERE userID=? ORDER BY timestamp DESC LIMIT ?, ?`
var query_GETLIKECOUNT = `SELECT COUNT(postID) FROM Like WHERE postID=?`
var query_GETCOMMENTCOUNT = `SELECT COUNT(postID) FROM Comment WHERE postID=?`

func (db *appdbimpl) GetPosts(profileUserID uint32, offset uint32, limit int32) ([]Post, error) {
	// Get the posts from the database
	rows, err := db.c.Query(query_GETPOSTS, profileUserID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer func() { err = rows.Close() }()

	// Create the slice of posts
	posts := make([]Post, 0)

	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}
		var post Post
		var user User
		// Get post data
		err = rows.Scan(&post.PostID, &user.UserID, &post.ImageURL, &post.Caption, &post.Timestamp)
		if err != nil {
			return nil, err
		}
		// Get like count
		err = db.c.QueryRow(query_GETLIKECOUNT, post.PostID).Scan(&post.LikeCount)
		if err != nil {
			return nil, err
		}

		// Get comment count
		err = db.c.QueryRow(query_GETCOMMENTCOUNT, post.PostID).Scan(&post.CommentCount)
		if err != nil {
			return nil, err
		}

		// Get user data
		user, err = db.GetUserByID(user.UserID)
		if err != nil {
			return nil, err
		}

		// Set user data
		post.User = user

		posts = append(posts, post)

	}
	return posts, err
}
