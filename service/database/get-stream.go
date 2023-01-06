package database

var join_FOLLOWINGS_POSTS = `SELECT User.userID, User.username, User.userPropicURL, Post.postID, Post.postImageURL, Post.caption, Post.timestamp 
							FROM (` + query_GETFOLLOWINGS + `) AS User INNER JOIN Post ON User.userID = Post.userID ORDER BY Post.timestamp DESC LIMIT ?, ?`

func (db *appdbimpl) GetStream(userID uint32, offeset uint32, limit int32) ([]Post, error) {
	var posts []Post

	// Get the posts from the database
	res, err := db.c.Query(join_FOLLOWINGS_POSTS, userID, 0, -1, offeset, limit)
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
		if err := res.Scan(&user.UserID, &user.Username, &user.UserPropicURL, &post.PostID, &post.ImageURL, &post.Caption, &post.Timestamp); err != nil {
			return nil, err
		}

		// Get the likes count for the post
		if err := db.c.QueryRow(query_GETLIKECOUNT, post.PostID).Scan(&post.LikeCount); err != nil {
			return nil, err
		}

		// Get the comments count for the post
		if err := db.c.QueryRow(query_GETCOMMENTCOUNT, post.PostID).Scan(&post.CommentCount); err != nil {
			return nil, err
		}

		// Set the user data
		post.User = user

		// Append the post to the posts slice
		posts = append(posts, post)
	}

	return posts, err
}
