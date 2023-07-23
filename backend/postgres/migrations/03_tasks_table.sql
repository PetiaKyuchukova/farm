-- +goose Up
CREATE TABLE tasks
(
    cowID text NOT NULL,
    date timestamp NOT NULL,
    type text NOT NULL,
    text text NOT NULL,
    done boolean,
    PRIMARY KEY (cowID,date),
    FOREIGN KEY (cowID) REFERENCES cows(id)

);

-- +goose Down
DROP TABLE tasks;
