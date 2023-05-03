-- +goose Up
create table users
(
    id bigserial primary key,
    name text not null,
    age int not null
);

-- +goose Down
drop table users;
