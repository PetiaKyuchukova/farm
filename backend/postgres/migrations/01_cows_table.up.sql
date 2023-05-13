-- +goose Up
CREATE TABLE cows
(
    id text NOT NULL,
    birthdate timestamp NOT NULL,
    farmerId text NOT NULL,
    colour text,
    motherId text,
    lastOvulation timestamp,
    lastBirth timestamp,
    isPregnant boolean,
    fertilization timestamp,
    givingBirthDate timestamp,

    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE cows;
