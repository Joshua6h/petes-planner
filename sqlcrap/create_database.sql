CREATE DATABASE petes_planner


CREATE TABLE users(
	user_id INT IDENTITY(1, 1) NOT NULL PRIMARY KEY,
	username VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(50) NOT NULL,
	email VARCHAR(200) NOT NULL,
	first_name VARCHAR(200) NOT NULL,
	last_name VARCHAR(200) NOT NULL
);


CREATE TABLE events(
	event_id INT IDENTITY(1, 1) NOT NULL PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
	description VARCHAR(500),
	start_datetime DATETIME NOT NULL,
	end_datetime DATETIME NOT NULL,
	group_id UNIQUEIDENTIFIER
);


CREATE TABLE user_events(
	user_event_id INT IDENTITY(1, 1) NOT NULL PRIMARY KEY,
	user_id INT FOREIGN KEY REFERENCES users(user_id) NOT NULL,
	event_id INT FOREIGN KEY REFERENCES events(event_id) NOT NULL
);

CREATE TABLE friends(
	friend_id INT IDENTITY(1, 1) NOT NULL PRIMARY KEY,
	user_id INT FOREIGN KEY REFERENCES users(user_id) NOT NULL,
	friends_with_id INT FOREIGN KEY REFERENCES users(user_id) NOT NULL
);