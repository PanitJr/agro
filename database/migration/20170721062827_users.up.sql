CREATE TABLE IF NOT EXISTS users (
    id serial,
    username VARCHAR(128) UNIQUE,
    encrytped_password VARCHAR(128),
    email VARCHAR(128) UNIQUE,
    date_of_birth timestamp with time zone default null,
    PRIMARY KEY (id)
);

INSERT INTO public.users(username, encrytped_password,  email, date_of_birth) VALUES
('panit', 'ct26487gyoailetrhbao', 'panit@adsv.com', NOW());