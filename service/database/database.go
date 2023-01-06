/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type appdbimpl struct {
	c   *sql.DB
	ctx context.Context
}

type AppDatabase interface {
	// Create a new user
	CreateUser(u User) (User, error)

	// Delete a user
	DeleteUser(userID uint32) error

	// Get the user's stream. The list is ordered by timestamp, descending.
	GetStream(userID uint32, offset uint32, limit int32) ([]Post, error)

	// Change the username of a user
	ChangeUsername(userID uint32, newUsername string) error

	// Set/Change the bio of a user
	SetBio(userID uint32, bio string) error

	// Change the profile picture of a user
	// ChangeProfilePic(userID uint32, newProfilePicURL string) error

	// Get a user profile
	GetUserProfile(userID uint32) (Profile, error)

	// Get a user by its username
	GetUserByName(username string) (User, error)

	// Create a new post in the database
	CreatePost(p Post) (Post, error)

	// Delete a post from the database
	DeletePost(ownerID uint32, postID uint32) error

	// Get a list of post from a user. The list is ordered by timestamp, descending.
	// The list is limited to `limit` elements, and the first element is the `offset`-th element.
	GetPosts(profileUserID uint32, offset uint32, limit int32) ([]Post, error)

	// Get the last postID in the database
	GetLastPostID(userID uint32) (uint32, error)

	// Create a new like in the database. The user "userID" likes the post "postID".
	CreateLike(ownerID uint32, userID uint32, postID uint32) error

	// Delete a like from the database. The user "userID" unlikes the post "postID".
	DeleteLike(ownerID uint32, postID uint32, userID uint32) error

	// Get
	GetLikes(ownerID uint32, postID uint32, offset uint32, limit int32) ([]User, error)

	// Create a new comment in the database. The user "userID" comments the post "postID".
	CreateComment(ownerID uint32, userID uint32, postID uint32, text string) (Comment, error)

	// Delete a comment from the database. The user "ownerID" deletes the comment "commentID".
	DeleteComment(commentID uint32, ownerID uint32, postID uint32) error

	// Get all comments from a post. The list is ordered by timestamp, descending.
	GetComments(ownerID uint32, postID uint32, offset uint32, limit int32) ([]Comment, error)

	// Add a follow relationship between two users
	CreateFollow(followerID uint32, followedID uint32) error

	// Remove a follow relationship between two users
	DeleteFollow(followerID uint32, followedID uint32) error

	// Get all users followed by the user "followerID"
	GetFollowings(followerID uint32, offset uint32, limit int32) ([]User, error)

	// Get all users following the user "followedID"
	GetFollowers(followedID uint32, offset uint32, limit int32) ([]User, error)

	// Add a ban relationship between two users
	CreateBan(bannerID uint32, bannedID uint32) error

	// Remove a ban relationship between two users
	DeleteBan(bannerID uint32, bannedID uint32) error

	// Get all users banned from the user "bannerID"
	GetBans(bannerID uint32, offset uint32, limit int32) ([]User, error)

	// Get a user by its ID
	GetUserByID(userID uint32) (User, error)

	// Return true if the banned has been banned by the banner, false otherwise
	IsBanned(bannerID uint32, bannedID uint32) (bool, error)

	// Return true if the user is following the other user, false otherwise
	IsFollowing(followerID uint32, followedID uint32) (bool, error)

	SearchUsers(userID uint32, search string, from_follow bool, offset uint32, limit int32) ([]User, error)

	// Return true if a username already exists, false otherwise
	ExistsName(username string) (bool, error)

	Ping() error
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if tables exists. If not, the database is empty or not exists all tables, and we need to create them.
	var tableCount uint8
	err := db.QueryRow(`SELECT count(name) FROM sqlite_master WHERE type='table';`).Scan(&tableCount)
	if err != nil {
		return nil, fmt.Errorf("error checking if database is empty: %w", err)
	}
	// The tables are six in total, so if the count is less than 6, we need to create them.
	if tableCount != 6 {

		// ---CREATE USER TABLE----//
		_, err = db.Exec(sql_TABLEUSER)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(`INSERT INTO User (userID, username, userPropicURL)
						VALUES ("0", "NULL", "NULL");`)
		if err != nil {
			return nil, err
		}

		// ---CREATE POST TABLE----//
		_, err = db.Exec(sql_TABLEPOST)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// ---CREATE LIKE TABLE----//
		_, err = db.Exec(sql_TABLELIKE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// ---CREATE COMMENT TABLE----//
		_, err = db.Exec(sql_TABLECOMMENT)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// ---CREATE FOLLOW TABLE----//
		_, err = db.Exec(sql_TABLEFOLLOW)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// ---CREATE BAN TABLE----//
		_, err = db.Exec(sql_TABLEBAN)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	return &appdbimpl{
		c:   db,
		ctx: context.Background(),
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
