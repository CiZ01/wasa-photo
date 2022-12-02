package database

type User struct {
	UserID        uint32 `json:"userID"`
	Username      string `json:"username"`
	UserPropicURL string `json:"userPropicURL"`
}
