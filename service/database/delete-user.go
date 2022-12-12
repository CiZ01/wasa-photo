package database

var query_DELETEALLLIKE = `DELETE FROM Like WHERE userID = ?`
var query_DELETEALLCOMMENT = `DELETE FROM Comment WHERE userID = ?`
var query_DELETEALLFOLLOW = `DELETE FROM Follow WHERE userID = ? OR targetUserID = ?`
var query_DELETEALLBAN = `DELETE FROM Ban WHERE userID = ? OR targetUserID = ?`
var query_DELETEALLPOSTS = `DELETE FROM Post WHERE userID = ?`
var query_DELETEUSER = `DELETE FROM User WHERE userID = ?`

func (db *appdbimpl) DeleteUser(userID uint32) error {
	_, err := db.c.Exec(query_DELETEALLPOSTS, userID)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(query_DELETEALLBAN, userID, userID)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(query_DELETEALLFOLLOW, userID, userID)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(query_DELETEALLCOMMENT, userID)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(query_DELETEALLLIKE, userID)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(query_DELETEUSER, userID)
	if err != nil {
		return err
	}
	return nil
}
