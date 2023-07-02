-- +goose Up
CREATE TABLE tasks
(
    cowID text NOT NULL,
    date timestamp NOT NULL,
    type text NOT NULL,
    text text NOT NULL,
    PRIMARY KEY (cowID,date)
);

-- +goose Down
DROP TABLE tasks;
