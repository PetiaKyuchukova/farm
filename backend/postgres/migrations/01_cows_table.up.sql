-- +goose Up
CREATE TABLE cows
(
    id text NOT NULL,
    birthdate timestamp,
    gender text,
    breed text,
    colour text,
    motherId text,
    motherBreed text,
    fatherId text,
    fatherBreed text,
    isPregnant boolean,
    ovulation timestamp,

    PRIMARY KEY (id)

);

-- +goose Down
DROP TABLE cows;
