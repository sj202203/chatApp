-- migrate:up
CREATE TABLE IF NOT EXISTS "user" (
    user_id BIGINT PRIMARY KEY,
    name varchar(100)
);

CREATE TABLE IF NOT EXISTS user_group (
    id BIGINT PRIMARY KEY,
    users JSONB
);

-- migrate:down
DROP TABLE IF EXISTS "user";

DROP TABLE IF EXISTS user_group;