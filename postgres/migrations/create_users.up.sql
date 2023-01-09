CREATE TABLE users(
    userid BIGSERIAL PRIMARY KEY,
    fname VARCHAR(55) NOT NULL,
    lname VARCHAR(55) NOT NULL,
    password TEXT,
    carid VARCHAR(55) NOT NULL
);