create table product (
    id serial PRIMARY KEY,
    name varchar(40) NOT NULL UNIQUE,
    price numeric(5, 2),
    since date
);

create table users (
    id serial PRIMARY KEY,
    name varchar
);

-- drop table product;
-- alter table product ADD COLUMN ()

