CREATE TABLE users (
    uid integer primary key autoincrement,
    name varchar(30),
    surname varchar(30),
    username varchar(30),
    email varchar(50),
    passwd_h varchar(50),
    reg_date timestamp,
    last_login timestamp, 
    cicerone boolean
);