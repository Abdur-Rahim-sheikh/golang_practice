-- +migrate Up


CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    title varchar(100) not null,
    description TEXT,
    price double precision not null,
    img_url TEXT,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp
);