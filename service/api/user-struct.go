package api

import (
	"regexp"

	"git.francescofazzari.it/wasa_photo/service/api/utils"
	"git.francescofazzari.it/wasa_photo/service/database"
)

/*
This struct rappresents the User object.
The user is identified by the username and the userId, which is the primary key.
*/
type User struct {
	UserID       int    `json:"userID"`
	Username     string `json:"username"`
	UserPropic64 string `json:"userPropic64"`
}

func (u *User) ToDatabase() database.User {
	return database.User{
		UserID:        u.UserID,
		Username:      u.Username,
	}
}

func (u *User) FromDatabase(dbUser database.User) error {
	u.UserID = dbUser.UserID
	u.Username = dbUser.Username
	propic64, err := utils.ImageToBase64(utils.GetProfilePicPath(u.UserID))
	u.UserPropic64 = propic64
	if err != nil {
		return err
	}
	return nil
}

func IsValid(username string) bool {
	validUser := regexp.MustCompile(`^[a-z][a-z0-9_]{2,13}$`)
	return validUser.MatchString(username)
}
