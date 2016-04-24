-- +migrate Up
CREATE TABLE if not exists blood_pressures (
    patient_id integer NOT NULL,
    value character varying(255),

    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE INDEX idx_blood_pressures_on_patient_id_and_created_at ON blood_pressures (patient_id, created_at);

-- +migrate Down
drop table blood_pressures;
drop index idx_blood_pressures_on_patient_id_and_created_at;