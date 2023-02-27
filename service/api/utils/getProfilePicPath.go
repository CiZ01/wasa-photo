package utils

import (
	"fmt"
)

func GetProfilePicPath(userID int) string {
	return fmt.Sprintf("./storage/%d/user_propic_250x250.jpg", userID)
}
