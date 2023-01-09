CREATE TABLE "user"(
    "id"        BIGSERIAL PRIMARY KEY,
    "username"  VARCHAR     NOT NULL,
    "email"     VARCHAR     NOT NULL,
    "password"  VARCHAR(70)     NOT NULL
)