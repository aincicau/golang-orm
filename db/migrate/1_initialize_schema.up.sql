BEGIN;

CREATE TABLE vehicles (
    id varchar(10) NOT NULL,
    name varchar(100) NOT NULL,
    vehicleType varchar(10) NULL,
    CONSTRAINT vehicle_key PRIMARY KEY (id)
);

CREATE TABLE persons (
    id varchar(10) NOT NULL,
    name varchar(100) NOT NULL,
    age INT NULL,
    CONSTRAINT person_key PRIMARY KEY (id)
);

COMMIT;