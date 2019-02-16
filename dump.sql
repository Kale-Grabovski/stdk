CREATE TABLE module (
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar (100) NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp
) ENGINE=InnoDB;

INSERT INTO module (name) VALUES ('proxy'), ('ads');

CREATE TABLE module_version (
    id int AUTO_INCREMENT PRIMARY KEY,
    module_id int,
    filename varchar(64) NOT NULL,
    filehash varchar(32) NOT NULL,
    settings json NOT NULL,
    is_active tinyint(1) DEFAULT FALSE,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    CONSTRAINT FOREIGN KEY (module_id) REFERENCES module(id)
) ENGINE=InnoDB;
