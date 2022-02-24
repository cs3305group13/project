DROP TABLE user_credentials;

CREATE TABLE user_credentials (
    id INT NOT NULL AUTO_INCREMENT,
    username text NOT NULL,
    hash text NOT NULL,
    salt text NOT NULL,
    PRIMARY KEY (id)
);

SELECT *
FROM user_credentials;