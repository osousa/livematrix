CREATE TABLE plexerbot.`Session` (
	id INT NOT NULL AUTO_INCREMENT,
	`session` varchar(100) NOT NULL,
	expirity varchar(100) NULL,
	alias varchar(100) NULL,
	email varchar(100) NULL,
	ip varchar(100) NULL,
	CONSTRAINT Session_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;


GRANT Alter ON plexerbot.* TO 'osousa'@'localhost';
GRANT Create ON plexerbot.* TO 'osousa'@'localhost';
GRANT Create view ON plexerbot.* TO 'osousa'@'localhost';
GRANT Delete ON plexerbot.* TO 'osousa'@'localhost';
GRANT Delete history ON plexerbot.* TO 'osousa'@'localhost';
GRANT Drop ON plexerbot.* TO 'osousa'@'localhost';
GRANT Grant option ON plexerbot.* TO 'osousa'@'localhost';
GRANT Index ON plexerbot.* TO 'osousa'@'localhost';
GRANT Insert ON plexerbot.* TO 'osousa'@'localhost';
GRANT References ON plexerbot.* TO 'osousa'@'localhost';
GRANT Select ON plexerbot.* TO 'osousa'@'localhost';
GRANT Show view ON plexerbot.* TO 'osousa'@'localhost';
FLUSH PRIVILEGES;

