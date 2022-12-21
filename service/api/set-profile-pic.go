package api

func (rt *_router) setMyProfilePic(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Get the user ID from the URL
	_profileUserID, err := strconv.Atoi(ps.ByName("profileUserID"))
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user is authorized
	if !isAuthorized(uint32(_profileUserID), r.Header) {
		rt.writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	// Get the file from the request
	file, _, err := r.FormFile("profilePic")
	if err != nil {
		http.Error(w, "Bad Request"+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get the file size
	fileStat, err := file.Stat()
	if err != nil {
		rt.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Check if the file size is greater than 5MB
	if fileStat.Size() > 5*1024*1024 {
		rt.writeError(w, http.StatusBadRequest, "The file size is too big")
		return
	}

	// Read the file content
	fileBytes, err := ioutil