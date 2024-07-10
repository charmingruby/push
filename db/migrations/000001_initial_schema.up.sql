CREATE TABLE IF NOT EXISTS examples
(
    id varchar PRIMARY KEY NOT NULL,
    name varchar NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL
);