-- +goose Up
CREATE TABLE pregnancies
(
    cowId text,
    detectedAt timestamp,
    firstDay timestamp,
    lastDay timestamp,
    PRIMARY KEY (cowId, detectedAt),
    FOREIGN KEY (cowId) REFERENCES cows(id)

);

-- +goose Down
DROP TABLE pregnancies;