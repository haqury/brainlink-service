-- +goose Up
-- +goose StatementBegin
CREATE TABLE brainlink.meta_eeg
(
    id         bigint PRIMARY KEY,
    eeg_history_id bigint,

    count      int default 0 NOT NULL
);

create sequence brainlink.meta_eeg_id_seq;

COMMENT ON TABLE brainlink.meta_eeg IS 'Метаданные еег';

ALTER TABLE ONLY brainlink.meta_eeg
    ADD CONSTRAINT meta_eeg_pkey
        PRIMARY KEY (id);

ALTER TABLE ONLY brainlink.meta_eeg
    ADD CONSTRAINT brainlink_fk_eeg_history_id
        FOREIGN KEY (eeg_history_id) REFERENCES brainlink.eeg_history (id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
