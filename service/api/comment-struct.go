package api

import "time"

/*
This struct rappresents the Comment object.
The post is identified by the CommentID, which is the primary key.
*/
type Comment struct {
	CommentID uint      `json:"commentID"`
	User      User      `json:"user"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
	// non so se sia il massimo usare time visto che
	// il timestamp lo gestirò lato frontend e comunque
	// quando gli arriva è una stringa
}
