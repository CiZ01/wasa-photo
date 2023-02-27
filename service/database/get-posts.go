package database

var query_GETPOSTS = `SELECT postID, userID, postImageURL, caption, timestamp FROM Post WHERE userID=? ORDER BY timestamp DESC LIMIT ?, ?`
var query_GETLIKECOUNT = `SELECT COUNT(postID) FROM Like WHERE postID=? AND ownerID=?`
var query_GETCOMMENTCOUNT = `SELECT COUNT(postID) FROM Comment WHERE postID=? AND ownerID=?`
var query_ISLIKED = `SELECT COUNT(postID) FROM Like WHERE postID=? AND ownerID=? AND userID=?`

func (db *appdbimpl) GetPosts(userID int, profileUserID int, offset int, limit int) ([]Post, error) {
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
		err = db.c.QueryRow(query_GETLIKECOUNT, post.PostID, profileUserID).Scan(&post.LikeCount)
		if err != nil {
			return nil, err
		}

		// Get comment count
		err = db.c.QueryRow(query_GETCOMMENTCOUNT, post.PostID, profileUserID).Scan(&post.CommentCount)
		if err != nil {
			return nil, err
		}

		// Get like status
		var like int
		err = db.c.QueryRow(query_ISLIKED, post.PostID, user.UserID, userID).Scan(&like)
		if err != nil {
			return nil, err
		}
		if like == 1 {
			post.Liked = true
		} else {
			post.Liked = false
		}

		// Get owner data
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
