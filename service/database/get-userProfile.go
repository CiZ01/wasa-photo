package database

var query_GETUSERINFO = `SELECT userID, username, bio FROM User WHERE userID=?;`
var query_GETCOUNTFOLLOWINGS = `SELECT count(followedID) FROM Follow WHERE followerID=?;`
var query_GETCOUNTFOLLOWERS = `SELECT count(followerID) FROM Follow WHERE followedID=?;`
var query_ISFOLLOWED = `SELECT count(followedID) FROM Follow WHERE followedID=? AND followerID=?;`

func (db *appdbimpl) GetUserProfile(profileUserID int, userID int) (Profile, error) {
	var profile Profile

	rows, err := db.c.Query(query_GETUSERINFO, profileUserID)
	if err != nil {
		return profile, err
	}
	defer func() { err = rows.Close() }()

	for rows.Next() {
		if rows.Err() != nil {
			return profile, err
		}
		var user User
		err := rows.Scan(&user.UserID, &user.Username, &profile.Bio)
		if err != nil {
			return profile, err
		}
		profile.User = user
	}

	err = db.c.QueryRow(query_GETCOUNTFOLLOWINGS, profileUserID).Scan(&profile.FollowingsCount)
	if err != nil {
		return profile, err
	}

	err = db.c.QueryRow(query_GETCOUNTFOLLOWERS, profileUserID).Scan(&profile.FollowersCount)
	if err != nil {
		return profile, err
	}

	err = db.c.QueryRow(query_ISFOLLOWED, profileUserID, userID).Scan(&profile.IsFollowed)
	if err != nil {
		return profile, err
	}

	return profile, err
}
