package api

import (
	"regexp"

	"git.francescofazzari.it/wasa_photo/service/database"
)

/*
This struct rappresents the User object.
The user is identified by the username and the userId, which is the primary key.
*/
type User struct {
	UserID        uint32 `json:"userID"`
	Username      string `json:"username"`
	UserPropicURL string `json:"userPropicURL"`
}

func (u *User) ToDatabase() database.User {
	return database.User{
		UserID:        u.UserID,
		Username:      u.Username,
		UserPropicURL: u.UserPropicURL,
	}
}

func (u *User) FromDatabase(User database.User) {
	u.UserID = User.UserID
	u.Username = User.Username
	u.UserPropicURL = User.UserPropicURL
}

func (u User) IsValid() bool {
	validUser := regexp.MustCompile(`^[a-z_]{1}[a-z0-9][a-z0-9_]{1,13}$`)
	return validUser.MatchString(u.Username)
}
