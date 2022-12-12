package database

// --------USER TABLE--------//
/*
	- userID: the unique ID of the user, is the primary key of the table
	- username: the username of the user
	- userPropicURL: the URL of the user's profile picture, it's point to the image in the storage.
	- bio: the bio of the user, is a short description of the user. It's optional. Max length is 64 characters.
	- timestamp: the timestamp of the user's creation. It's is assigned automatically by the database.
*/
var sql_TABLEUSER = `CREATE TABLE User
(
	userID INTEGER NOT NULL UNIQUE,
	username STRING NOT NULL UNIQUE,
	userPropicURL STRING NOT NULL,
	bio TEXT,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(userID)
);`

// --------POST TABLE--------//
/*
	- postID: the unique ID of the post, is the primary key of the table
	- userID: the userID of the user who created the post
	- postImageURL: the URL of the post's image, it's point to the image in the storage.
	- caption: the caption of the post, is a short description of the post. It's optional. Max length is 100 characters.
	- timestamp: the timestamp of the post's creation. It's is assigned automatically by the database.
*/
var sql_TABLEPOST = `CREATE TABLE Post
(
	postID INTEGER NOT NULL,
	userID INTEGER NOT NULL,
	postImageURL STRING NOT NULL UNIQUE,
	caption TEXT,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(postID, userID),
	FOREIGN KEY(userID) REFERENCES User(userID)
);`

// secondo me non serve salvarsi la preview in db, la previw avrà lo stesso nome del post originale
// e sarà salvata nella stessa cartella con un nome tipo URL280x280.jpg, quindi non serve salvarla in db

// --------LIKE TABLE--------//
/*
	- userID: the ID of the user who liked the post. It's the primary key of the table
	- postID: the ID of the post that the user liked. It's the primary key of the table
	- timestamp: the timestamp of the like's creation. It's is assigned automatically by the database.
*/
var sql_TABLELIKE = `CREATE TABLE Like
(
	userID INTEGER NOT NULL,
	postID INTEGER NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(userID, postID),
	FOREIGN KEY(userID) REFERENCES User(userID),
	FOREIGN KEY(postID) REFERENCES Post(postID)
);`

// --------COMMENT TABLE--------//
/*
	- commentID: the unique ID of the comment, is the primary key of the table
	- userID: the userID of the user who created the comment
	- postID: the postID of the post that the user commented
	- commentText: the text of the comment. Max length is 100 characters.
	- hidden: if the comment is hidden or not. Default value is false.
	- timestamp: the timestamp of the comment's creation. It's is assigned automatically by the database.
*/
var sql_TABLECOMMENT = `CREATE TABLE Comment
(
	commentID INTEGER NOT NULL UNIQUE,
	userID INTEGER NOT NULL,
	postID INTEGER NOT NULL,
	commentText TEXT NOT NULL,
	hidden BOOLEAN DEFAULT FALSE,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(commentID),
	FOREIGN KEY(userID) REFERENCES User(userID),
	FOREIGN KEY(postID) REFERENCES Post(postID)
);`

// --------FOLLOW TABLE--------//
/*
	- followerID: the userID of the user who is following the other user. It's the primary key of the table
	- followedID: the userID of the user who is followed by the other user. It's the primary key of the table
	- timestamp: the timestamp of the follow's creation. It's is assigned automatically by the database.
*/
var sql_TABLEFOLLOW = `CREATE TABLE Follow
(
	followerID INTEGER NOT NULL,
	followedID INTEGER NOT NULL,
	PRIMARY KEY(followerID, followedID),
	FOREIGN KEY(followedID) REFERENCES User(userID),
	FOREIGN KEY(followerID) REFERENCES User(userID)
);`

// --------BAN TABLE--------//
/*
	- bannerID: the userID of the user who banned the other user. It's the primary key of the table
	- bannedID: the userID of the user who is banned. It's the primary key of the table
*/
var sql_TABLEBAN = `CREATE TABLE Ban
(
	bannerID INTEGER NOT NULL,
	bannedID INTEGER NOT NULL,
	PRIMARY KEY(bannedID, bannerID),
	FOREIGN KEY(bannedID) REFERENCES User(userID),
	FOREIGN KEY(bannerID) REFERENCES User(userID)
);`
