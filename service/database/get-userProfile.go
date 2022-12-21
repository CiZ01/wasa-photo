package database

var query_GETUSERINFO = `SELECT userID, username, bio, userPropicURL FROM User WHERE userID=?;`
var query_GETCOUNTFOLLOWINGS = `SELECT count(followedID) FROM Follow WHERE followerID=?;`
var query_GETCOUNTFOLLOWERS = `SELECT count(followerID) FROM Follow WHERE followedID=?;`

func (db *appdbimpl) GetUserProfile(userID uint32) (Profile, error) {
	var profile Profile

	rows, err := db.c.Query(query_GETUSERINFO, userID)
	if err != nil {
		return profile, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserID, &user.Username, &profile.Bio, &user.UserPropicURL)
		if err != nil {
			return profile, err
		}
		profile.User = user
	}

	err = db.c.QueryRow(query_GETCOUNTFOLLOWINGS, userID).Scan(&profile.FollowingsCount)
	if err != nil {
		return profile, err
	}

	err = db.c.QueryRow(query_GETCOUNTFOLLOWERS, userID).Scan(&profile.FollowersCount)
	if err != nil {
		return profile, err
	}

	return profile, nil
}