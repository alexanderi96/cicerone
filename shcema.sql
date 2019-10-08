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

CREATE TABLE tours (
    tid integer primary key autoincrement,
    cid primary key references users(uid),
);

CREATE TABLE files(
    Title varchar(1000) not null,
    autoName varchar(255) not null,
    user_id references users(uid),
    created_date timestamp
);
