--
-- Create table products
--


CREATE TABLE products (
    id uuid DEFAULT gen_random_uuid(),
    name varchar(255) NOT NULL,
    price decimal NOT NULL,
    primary key (id)
);

ALTER TABLE products 
add column updated_at timestamp not null,
add column created_at timestamp not null;   