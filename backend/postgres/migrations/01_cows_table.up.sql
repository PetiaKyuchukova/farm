-- +goose Up
CREATE TABLE cows
(
    id text NOT NULL,
    birthdate timestamp NOT NULL,
    colour text,
    motherId text,
    farmerId text NOT NULL,
    lastOvulation timestamp,
    lastBirth timestamp,
    isPregnant boolean,
    fertilization timestamp,

    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE cows;
