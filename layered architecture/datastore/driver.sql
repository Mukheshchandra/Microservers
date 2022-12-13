DROP DATABASE IF EXISTS testingDB;
CREATE DATABASE testingDB;
USE testingDB;

CREATE TABLE Author(
    authorId int NOT NULL AUTO_INCREMENT,
	Firstname varchar(50),
    Lastname varchar(50),
    Dob varchar(50),
    Penname varchar(50),
    PRIMARY KEY(authorId));

CREATE TABLE Book(
    bookId int NOT NULL AUTO_INCREMENT,
    Title varchar(50),
	authorId int,
	Publication varchar(50),
	PublishedDate varchar(50),
	PRIMARY KEY(bookId),
	CONSTRAINT FK FOREIGN KEY(authorId) REFERENCES Author(authorId));
