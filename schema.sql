CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE IF EXISTS "records" CASCADE;
CREATE TABLE "records"
(
    record_id SERIAL,
    record_name character varying(500)  NOT NULL,
    record_genre character varying(50) NOT NULL,
    record_year character varying(4) NOT NULL,
    CONSTRAINT records_pkey PRIMARY KEY (record_id)
);

DROP TABLE IF EXISTS "artist" CASCADE;
CREATE TABLE "artist"
(
    artist_id SERIAL,
    artist_name character varying(150)  NOT NULL,
    CONSTRAINT artist_pkey PRIMARY KEY (artist_id),
    CONSTRAINT uq_artist_name UNIQUE (artist_name)
);

INSERT INTO artist(artist_name) VALUES('Ajda Pekkan');

DROP TABLE IF EXISTS "artist_records" CASCADE;
CREATE TABLE artist_records
(
    artist_id integer,
    record_id integer,
    FOREIGN KEY ("artist_id") REFERENCES "artist" ("artist_id")
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY ("record_id") REFERENCES "records" ("record_id")
        ON DELETE CASCADE
        ON UPDATE CASCADE
);



DROP TABLE IF EXISTS "users" CASCADE;
CREATE TABLE "users"
(
    user_id SERIAL,
    user_email character varying(128) NOT NULL,
    user_password character varying(128) NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (user_id),
    CONSTRAINT user_email_uq UNIQUE (user_email)
);

INSERT INTO users(user_email, user_password) VALUES('eylulkadioglu99@gmail.com', '5b0493439c6b3777db3a675b9a013a5c1b3d445fc095755856a51f388fdf3685');

DROP TABLE IF EXISTS "password_codes" CASCADE;
CREATE TABLE password_codes
(
    user_id integer,
    code character varying NOT NULL,
    CONSTRAINT password_codes_pkey PRIMARY KEY (code),
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
        ON DELETE CASCADE
        ON UPDATE CASCADE
);


/* CREATE TABLE "organizations" 
(
    "organization_id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "name" VARCHAR(255) NOT NULL,
    "identifier" VARCHAR(32) NOT NULL,
    "owner_email" VARCHAR(255) NOT NULL,
    "url" VARCHAR(255),
    "comments" TEXT,
    "created_at" BIGINT DEFAULT EXTRACT(EPOCH FROM now()) * 1000,
    "updated_at" BIGINT DEFAULT EXTRACT(EPOCH FROM now()) * 1000,

    PRIMARY KEY ("organization_id"),
    UNIQUE ("identifier")
);
 */