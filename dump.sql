CREATE TABLE module (
    id serial PRIMARY KEY,
    name text NOT NULL,
    created_at timestamp NOT NULL DEFAULT now()
);

INSERT INTO module (name) VALUES ('proxy'), ('ads');

CREATE TABLE module_version (
    id serial PRIMARY KEY,
    module_id int REFERENCES module(id),
    filename text NOT NULL,
    filehash text NOT NULL,
    settings json NOT NULL,
    is_active boolean DEFAULT FALSE,
    created_at timestamp NOT NULL DEFAULT now()
);
