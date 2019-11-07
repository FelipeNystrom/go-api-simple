CREATE DATABASE goapi;

CREATE USER goapi_usr WITH ENCRYPTED PASSWORD 'goapi!234';
CREATE USER goapi_mngr WITH ENCRYPTED PASSWORD 'goapi!!!34';
CREATE USER goapi_admin WITH ENCRYPTED PASSWORD 'goapi!!!!!!4';

GRANT goapi_usr TO goapi_mngr;
GRANT goapi_mngr TO goapi_admin;

REVOKE ALL ON DATABASE goapi FROM public;

GRANT CONNECT ON DATABASE goapi TO goapi_usr;

\connect goapi

CREATE SCHEMA posts_schema AUTHORIZATION goapi_admin;



SET search_path = posts_schema;


ALTER ROLE goapi_usr   IN DATABASE goapi SET search_path = posts_schema;
ALTER ROLE goapi_mngr   IN DATABASE goapi SET search_path = posts_schema;
ALTER ROLE goapi_admin   IN DATABASE goapi SET search_path = posts_schema;

GRANT USAGE ON SCHEMA posts_schema TO goapi_usr;
GRANT CREATE ON SCHEMA posts_schema TO goapi_admin;

ALTER DEFAULT PRIVILEGES FOR ROLE goapi_admin
GRANT SELECT ON TABLES TO goapi_usr;  -- only read

ALTER DEFAULT PRIVILEGES FOR ROLE goapi_admin
GRANT INSERT, UPDATE, DELETE, TRUNCATE ON TABLES TO goapi_mngr;  -- + write, TRUNCATE optional

ALTER DEFAULT PRIVILEGES FOR ROLE goapi_admin
GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO goapi_mngr;  -- SELECT, UPDATE are optional 

\connect goapi goapi_admin

CREATE TABLE posts (
    id serial PRIMARY KEY, 
    title varchar(255) UNIQUE, 
    body text, 
    author varchar(255)
    );


INSERT INTO posts (title, body, author) VALUES('hej', 'hej', 'En författare'); 

