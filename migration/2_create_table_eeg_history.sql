-- +goose Up
-- +goose StatementBegin
CREATE TABLE brainlink.eeg_history
(
    id         bigint NOT NULL,
    Attention  INTEGER, -- an integer representing the value of the attention level
    Meditation INTEGER, -- an integer representing the value of the meditation level
    Signal     INTEGER, -- an integer representing the value of the signal level
    Delta      INTEGER, -- an integer representing the value of the delta level
    Theta      INTEGER, -- an integer representing the value of the theta level
    LowAlpha   INTEGER, -- an integer representing the value of the low alpha level
    HighAlpha  INTEGER, -- an integer representing the value of the high alpha level
    LowBeta    INTEGER, -- an integer representing the value of the low beta level
    HighBeta   INTEGER, -- an integer representing the value of the high beta level
    LowGamma   INTEGER, -- an integer representing the value of the low gamma level
    HighGamma  INTEGER,  -- an integer representing the value of the high gamma level
    system_mouse_id bigint NOT NULL
);

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
