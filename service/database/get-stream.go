package database

var join_FOLLOWINGS_POSTS = `SELECT User.userID, User.username, Post.postID, Post.postImageURL, Post.caption, Post.timestamp 
							FROM (` + union_FOLLOWINGS_ME + `) AS User INNER JOIN Post ON User.userID = Post.userID ORDER BY Post.timestamp DESC LIMIT ?, ?`
var union_FOLLOWINGS_ME = query_GETFOLLOWINGS + ` UNION SELECT userID, username FROM User WHERE userID=?`

func (db *appdbimpl) GetStream(userID int, offeset int, limit int) ([]Post, error) {
	var posts []Post

	// Get the posts from the database
	res, err := db.c.Query(join_FOLLOWINGS_POSTS, userID, 0, -1, userID, offeset, limit)
	if err != nil {
		return nil, err
	}
	defer func() { err = res.Close() }()

	for res.Next() {
		if err := res.Err(); err != nil {
			return nil, err
		}
		var user User
		var post Post

		// Scan the result into the post struct
		if err := res.Scan(&user.UserID, &user.Username, &post.PostID, &post.ImageURL, &post.Caption, &post.Timestamp); err != nil {
			return nil, err
		}

		// Get the likes count for the post
		if err := db.c.QueryRow(query_GETLIKECOUNT, post.PostID, user.UserID).Scan(&post.LikeCount); err != nil {
			return nil, err
		}

		// Get the comments count for the post
		if err := db.c.QueryRow(query_GETCOMMENTCOUNT, post.PostID, user.UserID).Scan(&post.CommentCount); err != nil {
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

		// Set the user data
		post.User = user

		// Append the post to the posts slice
		posts = append(posts, post)
	}

	return posts, err
}
