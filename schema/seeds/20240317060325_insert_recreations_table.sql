-- +goose Up
-- +goose StatementBegin
INSERT INTO recreations (
    user_id,
    recreation_id,
    genre,
    title, content,
    target_number, required_time,
    youtube_id,
    publish, published_at,
    drafted_at
) VALUES
    (
        'c2cc015d-5753-4132-86fe-624787ae49df',
        'c5c29ef8-dd34-45f8-8013-e80840356ea4',
        '[1, 2]',
        'Example Title 1', 'Content Example 1',
        10, 30,
        NULL,
        TRUE, NULL,
        NULL
    ),
    (
        '87cab8ea-2252-4a4d-bc15-3fe6192bbf5d',
        '0c2643c1-4b20-40c6-9163-1d8567954334',
        '[3, 4]',
        'Example Title 2', 'Content Example 2',
        20, 40,
        NULL,
        TRUE, NULL,
        NULL
    ),
    (
        'c2cc015d-5753-4132-86fe-624787ae49df',
        'df3ed291-32b3-44d3-86c2-e01537e066c6',
        '[5, 6]',
        'Example Title 3', 'Content Example 3',
        30, 50,
        NULL,
        TRUE, NULL,
        NULL
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM recreations
WHERE
    recreation_id = 'c5c29ef8-dd34-45f8-8013-e80840356ea4'
    OR recreation_id = '0c2643c1-4b20-40c6-9163-1d8567954334'
    OR recreation_id = 'df3ed291-32b3-44d3-86c2-e01537e066c6';
-- +goose StatementEnd
