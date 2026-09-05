CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    email varchar(255) unique not null,
    password varchar(255) not null,
    is_shop_owner boolean default FALSE,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title varchar(100) not null,
    description TEXT,
    price double precision not null,
    img_url TEXT,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp
);