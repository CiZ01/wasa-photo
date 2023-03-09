package api

/*
CreateUser function is used to create a new user.
It takes the user object from the doLogin function and returns the user object with
the userID
*/
func (rt *_router) CreateUser(u User) (User, error) {
	// Return the user object with the userID
	dbUser, err := rt.db.CreateUser(u.ToDatabase())
	if err != nil {
		return u, err
	}

	// Convert the database user object to the user object and return it.
	err = u.FromDatabase(dbUser)
	if err != nil {
		return u, err
	}

	return u, nil

}
