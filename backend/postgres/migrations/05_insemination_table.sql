-- +goose Up
CREATE TABLE inseminations
(
    cowId text,
    date timestamp,
    breed text,
    isArtificial bool,
    PRIMARY KEY (cowId, date)
);

-- +goose Down
DROP TABLE inseminations;