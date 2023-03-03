package utils

import (
	"fmt"
)

func GetPostPhotoPath(ownerID int, postID int) string {
	return fmt.Sprintf("./storage/%d/posts/%d.jpeg", ownerID, postID)
}
