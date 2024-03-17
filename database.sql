CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    full_name varchar(50) not null ,
    email TEXT UNIQUE NOT NULL ,
    age INT NOT NULL,
    division varchar(20) not null
);

SELECT * FROM employees;

