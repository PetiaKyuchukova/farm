-- +goose Up
CREATE TABLE cows
(
    id text NOT NULL,
    birthdate timestamp NOT NULL,
    colour text,
    motherid text,
    PRIMARY KEY (id,birthdate)
);

-- +goose Down
DROP TABLE cows;
