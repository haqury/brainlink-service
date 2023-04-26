-- +goose Up
-- +goose StatementBegin
CREATE TABLE brainlink.event_eeg
(
    id         bigint default NOT NULL,
    eeg_history_ids jsonb,
    asset
);

create sequence brainlink.eeg_history_id_seq;

COMMENT ON TABLE brainlink.eeg_history IS 'Пакеты загрузки';

COMMENT ON Column brainlink.eeg_history.Attention IS 'an integer representing the value of the attention level';
COMMENT ON Column brainlink.eeg_history.Meditation IS 'an integer representing the value of the meditation level';
COMMENT ON Column brainlink.eeg_history.Signal IS 'an integer representing the value of the signal level';
COMMENT ON Column brainlink.eeg_history.Delta IS 'an integer representing the value of the delta level';
COMMENT ON Column brainlink.eeg_history.Theta IS 'an integer representing the value of the theta level';
COMMENT ON Column brainlink.eeg_history.LowAlpha IS 'an integer representing the value of the low alpha level';
COMMENT ON Column brainlink.eeg_history.HighAlpha IS 'an integer representing the value of the high alpha level';
COMMENT ON Column brainlink.eeg_history.LowBeta IS 'an integer representing the value of the low beta level';
COMMENT ON Column brainlink.eeg_history.HighBeta IS 'an integer representing the value of the high beta level';
COMMENT ON Column brainlink.eeg_history.LowGamma IS 'an integer representing the value of the low gamma level';
COMMENT ON Column brainlink.eeg_history.HighGamma IS 'an integer representing the value of the low HighGamma level';

ALTER TABLE ONLY brainlink.eeg_history
    ADD CONSTRAINT eeg_history_pkey
        PRIMARY KEY (id);

ALTER TABLE ONLY brainlink.eeg_history
    ADD CONSTRAINT packages_fk_status_id
        FOREIGN KEY (system_mouse_id) REFERENCES brainlink.system_mouse (id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
