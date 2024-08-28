CREATE TABLE users (
    id INT AUTO_INCREMENT,
    user_id int UNSIGNED,
    username varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now() on update now(),
    PRIMARY KEY (id)
);