CREATE TABLE profiles (
	id int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY,
	uuid uuid NOT NULL PRIMARY KEY,
	nickname text NOT NULL,
	icon_url text NULL,
	created_at timestamptz NULL DEFAULT now(),
	updated_at timestamptz NULL DEFAULT now()
);

INSERT INTO profiles
    (nickname, uuid, icon_url)
VALUES
    ('suzurikawa', 'c2cc015d-5753-4132-86fe-624787ae49df', 'https://avatars.githubusercontent.com/u/1000000001');

INSERT INTO profiles
    (nickname, uuid, icon_url)
VALUES
    ('tanaka', '87cab8ea-2252-4a4d-bc15-3fe6192bbf5d', 'https://avatars.githubusercontent.com/u/1000000002');