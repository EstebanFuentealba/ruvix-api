-- +goose Up
-- +goose StatementBegin
CREATE extension
if not exists pgcrypto;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
drop extension pgcrypto;
-- +goose StatementEnd
