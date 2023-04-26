-- +goose Up
-- +goose StatementBegin
CREATE TABLE brainlink.system_mouse
(
    id bigint PRIMARY KEY,
    X INT,
    Y INT,
    ToX INT,
    ToY INT,
    EndX INT,
    EndY INT

);

create sequence brainlink.system_mouse_id_seq;

COMMENT ON TABLE brainlink.system_mouse IS 'mouse position';

COMMENT ON Column brainlink.system_mouse.X IS 'mouse the value of the x level';
COMMENT ON Column brainlink.system_mouse.Y IS 'mouse the value of the y level';
COMMENT ON Column brainlink.system_mouse.ToX IS 'mouse change x level';
COMMENT ON Column brainlink.system_mouse.ToY IS 'mouse change y level';
COMMENT ON Column brainlink.system_mouse.EndX IS 'mouse end x level';
COMMENT ON Column brainlink.system_mouse.EndY IS 'mouse end y level';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
