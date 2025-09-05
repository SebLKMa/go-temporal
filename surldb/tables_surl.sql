/*
Login to check your tables:
psql "host=localhost port=5432 dbname=<your_db> user=<username> password=<password>"
\dt
\d
*/

DROP TABLE IF EXISTS surl_id;
CREATE TABLE surl_id (
    rowid SERIAL PRIMARY KEY,
	unique_id TEXT,
	long_url TEXT,
    short_url TEXT,
    expires_on TIMESTAMPTZ NOT NULL,
	created_on TIMESTAMPTZ NOT NULL,
	modified_on TIMESTAMPTZ NOT NULL
);

DROP INDEX IF EXISTS surl_id_long_url_idx;
CREATE INDEX surl_id_long_url_idx ON surl_id (long_url);
DROP INDEX IF EXISTS surl_id_short_url_idx;
CREATE INDEX surl_id_short_url_idx ON surl_id (short_url);
