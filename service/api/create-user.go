package api

/*
CreateUser function is used to create a new user.
It takes the user object from the doLogin function and returns the user object with
the userID and the userPropicURL set.
*/
func (rt *_router) CreateUser(u User) (User, error) {
	// Return the user object with the userID and the userPropicURL set.
	dbUser, err := rt.db.CreateUser(u.ToDatabase())
	if err != nil {
		return u, err
	}

	// Convert the database user object to the user object and return it.
	u.FromDatabase(dbUser)

	return u, nil

}
