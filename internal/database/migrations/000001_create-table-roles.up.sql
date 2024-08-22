CREATE TABLE roles (
   id INT(64) primary key NOT NULL AUTO_INCREMENT,
   name TEXT NOT NULL,
   created_at TIMESTAMP default now()
);