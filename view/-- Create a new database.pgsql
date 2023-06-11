-- Create a new database
CREATE DATABASE db;


-- Create a new table
CREATE TABLE users (
    email VARCHAR(255) NOT NULL,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    PRIMARY KEY (email)
);