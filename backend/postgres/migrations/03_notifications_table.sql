-- +goose Up
CREATE TABLE notifications
(
    id text not null ,
    cowID text NOT NULL,
    date timestamp NOT NULL,
    type text NOT NULL,
    text text NOT NULL,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE notifications;
