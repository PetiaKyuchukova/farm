-- +goose Up
CREATE TABLE pregnancies
(
    cowId text,
    detectedAt timestamp,
    firstDay timestamp,
    lastDay timestamp,
    PRIMARY KEY (cowId, detectedAt)
);

-- +goose Down
DROP TABLE pregnancies;