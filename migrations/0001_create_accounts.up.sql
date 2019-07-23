BEGIN;

CREATE TABLE "accounts" (
    "id" CHAR(26) PRIMARY,
    "name" VARCHAR(255),
    "contact_email" VARCHAR(255)
);

CREATE INDEX "idx_accounts_name" ON "accounts" ("name");

COMMIT;
