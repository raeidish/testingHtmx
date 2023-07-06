CREATE TABLE IF NOT EXISTS todo (
    id int PRIMARY KEY,
    text varchar(255) NOT NULL,
    created timestamp
);
