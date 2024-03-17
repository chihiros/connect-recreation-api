-- +goose Up
-- +goose StatementBegin
INSERT INTO users
    (nickname, user_id, icon_url)
VALUES
    ('suzurikawa', 'c2cc015d-5753-4132-86fe-624787ae49df', 'https://avatars.githubusercontent.com/u/1000000001')
    ,('tanaka', '87cab8ea-2252-4a4d-bc15-3fe6192bbf5d', 'https://avatars.githubusercontent.com/u/1000000002')
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE
    user_id = 'c2cc015d-5753-4132-86fe-624787ae49df'
    OR user_id = '87cab8ea-2252-4a4d-bc15-3fe6192bbf5d'
;
-- +goose StatementEnd
