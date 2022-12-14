package api

/*
This struct rappresents the Profile object.
The profile is identified by the User.UserID, which is the primary key.
*/
type Profile struct {
	User            User   `json:"user"`
	UserPropicURL   string `json:"userPropicURL"`
	Bio             string `json:"bio"`
	FollowerCount   int    `json:"followerCount"`
	FollowingsCount int    `json:"followingsCount"`
}
