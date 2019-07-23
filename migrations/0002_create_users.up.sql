BEGIN;

CREATE TABLE "users" (
    "id" CHAR(26) PRIMARY,
    "account_id" CHAR(26) REFERENCES "accounts",
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "last_login" TIMESTAMP NULL,
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW
);

CREATE INDEX "idx_users_account_id" ON "users" ("account_id");
CREATE INDEX "idx_users_name" ON "users" ("name");
CREATE UNIQUE INDEX "uidx_users_name" ON "users" ("email");

COMMIT;
