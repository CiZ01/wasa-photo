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
	DeleteUser(userID int) error

	// Get the user's stream. The list is ordered by timestamp, descending.
	GetStream(userID int, offset int, limit int) ([]Post, error)

	// Change the username of a user
	ChangeUsername(userID int, newUsername string) error

	// Set/Change the bio of a user
	SetBio(userID int, bio string) error

	// Get a user profile
	GetUserProfile(profileUserID int, userID int) (Profile, error)

	// Get a user by its username
	GetUserByName(username string) (User, error)

	// Create a new post in the database
	CreatePost(p Post, data []byte) (Post, error)

	// Change the post's caption
	UpdateCaption(ownerID int, postID int, newCaption string) error

	// Delete a post from the database
	DeletePost(ownerID int, postID int) error

	// Get a list of post from a user. The list is ordered by timestamp, descending.
	// The list is limited to `limit` elements, and the first element is the `offset`-th element.
	GetPosts(userID int, profileUserID int, offset int, limit int) ([]Post, error)

	// Get the last postID in the database
	GetLastPostID(userID int) (int, error)

	// Create a new like in the database. The user "userID" likes the post "postID".
	CreateLike(ownerID int, userID int, postID int) error

	// Delete a like from the database. The user "userID" unlikes the post "postID".
	DeleteLike(ownerID int, userID int, postID int) error

	// Get
	GetLikes(ownerID int, userID int, postID int, limit int) ([]User, error)

	// Create a new comment in the database. The user "userID" comments the post "postID".
	CreateComment(ownerID int, userID int, postID int, text string) (Comment, error)

	// Delete a comment from the database. The user "ownerID" deletes the comment "commentID".
	DeleteComment(commentID int, ownerID int, postID int) error

	// Get all comments from a post. The list is ordered by timestamp, descending.
	GetComments(ownerID int, postID int, offset int, limit int) ([]Comment, error)

	// Add a follow relationship between two users
	CreateFollow(followerID int, followedID int) error

	// Remove a follow relationship between two users
	DeleteFollow(followerID int, followedID int) error

	// Get all users followed by the user "followerID"
	GetFollowings(followerID int, offset int, limit int) ([]User, error)

	// Get all users following the user "followedID"
	GetFollowers(followedID int, offset int, limit int) ([]User, error)

	// Add a ban relationship between two users
	CreateBan(bannerID int, bannedID int) error

	// Remove a ban relationship between two users
	DeleteBan(bannerID int, bannedID int) error

	// Get all users banned from the user "bannerID"
	GetBans(bannerID int, offset int, limit int) ([]User, error)

	// Get a user by its ID
	GetUserByID(userID int) (User, error)

	// Return true if the banned has been banned by the banner, false otherwise
	IsBanned(bannerID int, bannedID int) (bool, error)

	// Return true if the user is following the other user, false otherwise
	IsFollowing(followerID int, followedID int) (bool, error)

	SearchUsers(userID int, search string, offset int, limit int) ([]User, error)

	// Return true if a username already exists, false otherwise
	ExistsName(username string) (bool, error)

	// Ping the database to check if it is alive
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
