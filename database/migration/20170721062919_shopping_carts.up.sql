CREATE TABLE IF NOT EXISTS shopping_carts (
    id serial, 
    user_id integer REFERENCES users (id) ON DELETE RESTRICT,
    product_id integer REFERENCES products (id) ON DELETE RESTRICT,
    qountity integer, 
    unit VARCHAR(32), 
    status VARCHAR(32), 
    PRIMARY KEY (id)
);