-- +goose Up
CREATE TABLE farmers
(
    id text not null ,
    username text NOT NULL,
    name text NOT NULL,
    password text NOT NULL,
    email text NOT NULL,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE farmers;