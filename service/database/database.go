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
	"database/sql"
	"errors"
	"fmt"
)

type appdbimpl struct {
	c *sql.DB
}

type AppDatabase interface {
	// Create a new user
	CreateUser(u User) (User, error)

	// Change the username of a user
	ChangeUsername(userID uint32, newUsername string) error

	// Get a user by its username
	GetUserByName(username string) (User, error)

	// Add a follow relationship between two users
	CreateFollow(followerID uint32, followedID uint32) error

	// Get a user by its ID
	GetUserByID(userID uint32) (User, error)

	// Return true if the banned has been banned by the banner, false otherwise
	IsBanned(bannerID uint32, bannedID uint32) (bool, error)

	// Return true if the user is following the other user, false otherwise
	IsFollowing(followerID uint32, followedID uint32) (bool, error)

	// Return true if a username already exists, false otherwise
	ExistsName(username string) bool

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

		//----CREATE USER TABLE----//
		_, err = db.Exec(sql_TABLEUSER)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(`INSERT INTO User (userID, username, userPropicURL)
						VALUES ("0", "haru", "NULL");`)
		if err != nil {
			return nil, err
		}

		//----CREATE POST TABLE----//
		_, err = db.Exec(sql_TABLEPOST)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		//----CREATE LIKE TABLE----//
		_, err = db.Exec(sql_TABLELIKE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		//----CREATE COMMENT TABLE----//
		_, err = db.Exec(sql_TABLECOMMENT)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		//----CREATE FOLLOW TABLE----//
		_, err = db.Exec(sql_TABLEFOLLOW)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		//----CREATE BAN TABLE----//
		_, err = db.Exec(sql_TABLEBAN)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
