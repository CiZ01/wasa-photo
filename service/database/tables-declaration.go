package database

// --------USER TABLE--------//
/*
	- userID: the unique ID of the user, is the primary key of the table
	- username: the username of the user
	- bio: the bio of the user, is a short description of the user. It's optional. Max length is 64 characters.
	- timestamp: the timestamp of the user's creation. It's is assigned automatically by the database.
*/
var sql_TABLEUSER = `CREATE TABLE IF NOT EXISTS User
(
	userID INTEGER NOT NULL UNIQUE,
	username STRING NOT NULL UNIQUE,
	bio TEXT,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(userID)
);`

// --------POST TABLE--------//
/*
	- postID: the unique ID of the post, is the primary key of the table
	- userID: the userID of the user who created the post
	- caption: the caption of the post, is a short description of the post. It's optional. Max length is 100 characters.
	- timestamp: the timestamp of the post's creation. It's is assigned automatically by the database.
*/
var sql_TABLEPOST = `CREATE TABLE IF NOT EXISTS Post
(
	postID INTEGER NOT NULL,
	userID INTEGER NOT NULL,
	caption TEXT,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(postID, userID),
	CONSTRAINT fk_post
		FOREIGN KEY (userID) REFERENCES User(userID)
		ON DELETE CASCADE
);`

// secondo me non serve salvarsi la preview in db, la previw avrà lo stesso nome del post originale
// e sarà salvata nella stessa cartella con un nome tipo URL280x280.jpg, quindi non serve salvarla in db

// --------LIKE TABLE--------//
/*
	- userID: the ID of the user who liked the post. It's the primary key of the table
	- ownerID: the ID of the user who owns the post that the user liked.
	- postID: the ID of the post that the user liked. It's the primary key of the table
	- timestamp: the timestamp of the like's creation. It's is assigned automatically by the database.
*/
var sql_TABLELIKE = `CREATE TABLE IF NOT EXISTS Like
(
	userID INTEGER NOT NULL,
	ownerID INTEGER NOT NULL,
	postID INTEGER NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(userID, ownerID, postID),
	CONSTRAINT fk_like
		FOREIGN KEY (ownerID, postID) REFERENCES Post(userID, postID)
		FOREIGN KEY (userID) REFERENCES User(userID)
		ON DELETE CASCADE
);`

// --------COMMENT TABLE--------//
/*
	- commentID: the unique ID of the comment, is one of the primary keys of the table
	- userID: the userID of the user who created the comment.
	- ownerID: the userID of the user who owns the post that the user commented, is one of the primary keys of the table
	- postID: the postID of the post that the user commented, is one of the primary keys of the table
	- commentText: the text of the comment. Max length is 100 characters.
	- hidden: if the comment is hidden or not. Default value is false.
	- timestamp: the timestamp of the comment's creation. It's is assigned automatically by the database.
*/
var sql_TABLECOMMENT = `CREATE TABLE IF NOT EXISTS Comment
(
	commentID INTEGER NOT NULL,
	userID INTEGER NOT NULL,
	ownerID INTEGER NOT NULL,
	postID INTEGER NOT NULL,
	commentText TEXT NOT NULL,
	hidden BOOLEAN DEFAULT FALSE,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(commentID, ownerID, postID),
	CONSTRAINT fk_comment
		FOREIGN KEY (ownerID, postID) REFERENCES Post(userID, postID)
		FOREIGN KEY (userID) REFERENCES User(userID)
		ON DELETE CASCADE
);`

// --------FOLLOW TABLE--------//
/*
	- followerID: the userID of the user who is following the other user. It's the primary key of the table
	- followedID: the userID of the user who is followed by the other user. It's the primary key of the table
*/
var sql_TABLEFOLLOW = `CREATE TABLE IF NOT EXISTS Follow
(
	followerID INTEGER NOT NULL,
	followedID INTEGER NOT NULL,
	PRIMARY KEY(followerID, followedID),
	CONSTRAINT fk_follow
	FOREIGN KEY (followerID) REFERENCES User(userID)
	FOREIGN KEY (followerID) REFERENCES User(userID)
		ON DELETE CASCADE
);`

// --------BAN TABLE--------//
/*
	- bannerID: the userID of the user who banned the other user. It's the primary key of the table
	- bannedID: the userID of the user who is banned. It's the primary key of the table
*/
var sql_TABLEBAN = `CREATE TABLE IF NOT EXISTS Ban
(
	bannerID INTEGER NOT NULL,
	bannedID INTEGER NOT NULL,
	PRIMARY KEY(bannedID, bannerID),
	CONSTRAINT fk_ban
		FOREIGN KEY (bannerID) REFERENCES User(userID)
		FOREIGN KEY (bannedID) REFERENCES User(userID)
		ON DELETE CASCADE
);`
