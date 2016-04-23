-- +migrate Up
CREATE TABLE if not exists patients (
    id integer NOT NULL,
    firstname character varying(255),
    lastname character varying(255),
    email character varying(255) DEFAULT ''::character varying NOT NULL,
    encrypted_password character varying(255) DEFAULT ''::character varying NOT NULL,
    social_security_number character varying(12),
    birth_date date,
    gender character varying(255),
    address character varying(255),
    zipcode character varying(255),
    city character varying(255),
    country character varying(255),    
    heart_rate_min character varying(3),
    heart_rate_max character varying(3),
    weight character varying(3),
    doctor_id integer,

    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE UNIQUE INDEX idx_patients_on_social_security_number ON patients (social_security_number);
CREATE UNIQUE INDEX idx_patients_on_email ON users (email);

-- +migrate Down
drop table patients;
drop index idx_patients_on_social_security_number;
drop index idx_patients_on_email;