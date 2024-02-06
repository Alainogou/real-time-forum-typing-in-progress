CREATE TABLE user (
 id INTEGER NOT NULL PRIMARY KEY,
 nickName VARCHAR(30) UNIQUE NOT NULL,
 passWord VARCHAR(100) NOT NULL,
 email VARCHAR(30) UNIQUE NOT NULL,
 firstName VARCHAR(30) NOT NULL,
 lastName VARCHAR(30) NOT NULL,
 age INTEGER NOT NULL,
 gender VARCHAR(10) NOT NULL

);

CREATE TABLE category (
 id INTEGER NOT NULL PRIMARY KEY,
 categoryName VARCHAR(30) NOT NULL
);


CREATE TABLE post (
 id INTEGER NOT NULL PRIMARY KEY,
 userId INTEGER NOT NULL,
 title VARCHAR(30) NOT NULL,
 content TEXT NOT NULL,
 imgUrl VARCHAR(100),
 FOREIGN KEY(userId) REFERENCES user(id)
);

CREATE TABLE comment (
 id INTEGER NOT NULL PRIMARY KEY,
 userId INTEGER NOT NULL,
 postId INTEGER NOT NULL,
 content TEXT NOT NULL,
 created_at TIMESTAMP NOT NULL,
 FOREIGN KEY(userId) REFERENCES user(id),
 FOREIGN KEY(postId) REFERENCES post(id)
);

CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY NOT NULL,
    sessionId varchar(250),
    email VARCHAR(250)  NOT NULL,
    expires_at TIMESTAMP NOT NULL
);

CREATE TABLE user_post_reaction (
 id INTEGER NOT NULL PRIMARY KEY,
 userId INTEGER NOT NULL,
 postId INTEGER NOT NULL,
 isLiked TINYINT(1) NOT NULL,
 FOREIGN KEY(userId) REFERENCES user(id),
 FOREIGN KEY(postId) REFERENCES post(id)
);

CREATE TABLE user_comment_reaction (
 id INTEGER NOT NULL PRIMARY KEY,
 userId INTEGER NOT NULL,
 commentId INTEGER NOT NULL,
 isLiked TINYINT(1) NOT NULL,
 FOREIGN KEY(userId) REFERENCES user(id),
 FOREIGN KEY(commentId) REFERENCES comment(id)
);

CREATE TABLE category_relation (
 id INTEGER NOT NULL PRIMARY KEY,
 categoryId INTEGER NOT NULL,
 postId INTEGER NOT NULL,
 FOREIGN KEY( categoryId) REFERENCES category(id),
 FOREIGN KEY(postId) REFERENCES post(id)
);

CREATE TABLE message (
 id INTEGER NOT NULL PRIMARY KEY,
 fromUser INTEGER NOT NULL,
 toUser INTEGER NOT NULL,
 isRead TINYINT(1) NOT NULL,
 message TEXT NOT NULL,
 createdDate DATETIME NOT NULL,
 FOREIGN KEY(fromUser) REFERENCES user(id),
 FOREIGN KEY(toUser) REFERENCES user(id)
);


INSERT INTO user (id, nickName, passWord, email, firstName, lastName, age, gender)
VALUES
    (1, 'aladji', 'Mvlick@123', 'ass.lo@gmail.com', 'aladji', 'malick', 99, 'male');

INSERT INTO category (categoryName) VALUES ('Sport'), ('Art'), ('Informatics'), ('Religion'), ('Game');