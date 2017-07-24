CREATE TABLE IF NOT EXISTS products (
    id serial,
    user_id integer REFERENCES users (id) ON DELETE RESTRICT,
    name VARCHAR(128),
    unit_price decimal(10,2),
    description text,
    date_of_birth timestamp with time zone default null,
    PRIMARY KEY (id)
);