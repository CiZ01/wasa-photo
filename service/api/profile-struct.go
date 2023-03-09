package api

import (
	"git.francescofazzari.it/wasa_photo/service/database"
)

/*
This struct rappresents the Profile object.
The profile is identified by the User.UserID, which is the primary key.
*/
type Profile struct {
	User            User   `json:"user"`
	Bio             string `json:"bio"`
	FollowerCount   int    `json:"followersCount"`
	FollowingsCount int    `json:"followingsCount"`
	PostsCount      int    `json:"postsCount"`
	IsFollowed      bool   `json:"isFollowed"`
}

func (p *Profile) FromDatabase(dbProfile database.Profile) error {
	var apiUser User
	err := apiUser.FromDatabase(dbProfile.User)
	if err != nil {
		return err
	}

	p.User = apiUser
	p.Bio = dbProfile.Bio
	p.FollowerCount = dbProfile.FollowersCount
	p.FollowingsCount = dbProfile.FollowingsCount
	p.PostsCount = dbProfile.PostsCount

	if dbProfile.IsFollowed == 1 {
		p.IsFollowed = true
	} else {
		p.IsFollowed = false
	}

	return nil
}
