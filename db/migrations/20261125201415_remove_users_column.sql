-- migrate:up
ALTER TABLE IF EXISTS user_group
    DROP COLUMN IF EXISTS users;

-- migrate:down
ALTER TABLE IF EXISTS user_group
    ADD COLUMN IF NOT EXISTS users JSONB;