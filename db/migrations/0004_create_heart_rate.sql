-- +migrate Up
CREATE TABLE if not exists heart_rates (
    patient_id integer NOT NULL,
    value character varying(255),

    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE INDEX idx_heart_rates_on_patient_id_and_created_at ON heart_rates (patient_id, created_at);

-- +migrate Down
drop table heart_rate;
drop index idx_heart_rates_on_patient_id_and_created_at;