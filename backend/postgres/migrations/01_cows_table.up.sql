-- +goose Up
CREATE TABLE cows
(
    id text NOT NULL,
    birthdate timestamp NOT NULL,
    gender text,
    breed text,
    colour text,
    motherId text,
    fatherId text,
    fatherBreed text,
    isPregnant boolean,

    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE cows;
