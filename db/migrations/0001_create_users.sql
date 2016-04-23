-- +migrate Up
CREATE TABLE if not exists users (
    id integer NOT NULL,
    email character varying(255) DEFAULT ''::character varying NOT NULL,
    password character varying(255) DEFAULT ''::character varying NOT NULL,
    firstname character varying(255) DEFAULT ''::character varying,
    lastname character varying(255) DEFAULT ''::character varying,
   
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE UNIQUE INDEX idx_users_on_email ON users (email);

-- +migrate Down
drop table users;
drop index idx_users_on_email;