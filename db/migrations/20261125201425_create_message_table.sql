-- migrate:up
CREATE TABLE IF NOT EXISTS "message" (
    id BIGSERIAL PRIMARY KEY,
    receiver_id BIGINT,
    message VARCHAR(1000),
    CONSTRAINT "fk_user_message" FOREIGN KEY ("receiver_id") REFERENCES "user"("user_id")
);

-- migrate:down

DROP TABLE "message";