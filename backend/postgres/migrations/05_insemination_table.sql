-- +goose Up
CREATE TABLE inseminations
(
    cowId text,
    date timestamp,
    breed text,
    isArtificial bool,
    PRIMARY KEY (cowId, date),
    FOREIGN KEY (cowId) REFERENCES cows(id)
);

-- +goose Down
DROP TABLE inseminations;