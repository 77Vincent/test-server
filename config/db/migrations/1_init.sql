create schema if not exists main;
ALTER DEFAULT PRIVILEGES IN SCHEMA main GRANT ALL PRIVILEGES ON TABLES TO PUBLIC;
ALTER DEFAULT PRIVILEGES IN SCHEMA main GRANT ALL PRIVILEGES ON TYPES TO PUBLIC;
ALTER DEFAULT PRIVILEGES IN SCHEMA main GRANT ALL PRIVILEGES ON ROUTINES TO PUBLIC;

set search_path to main,public;

create table if not exists event (id serial);