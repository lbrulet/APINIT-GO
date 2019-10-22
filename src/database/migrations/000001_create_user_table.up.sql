BEGIN;

CREATE TABLE users
(
	id int auto_increment
		primary key,
	username varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	admin boolean DEFAULT false,
	verified boolean DEFAULT false,
	auth_method int NOT NULL,
	password varchar(255) NOT NULL
);

COMMIT;