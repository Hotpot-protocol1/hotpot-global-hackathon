-- +migrate Up

CREATE TABLE IF NOT EXISTS hotpot_user
(
    id SERIAL PRIMARY KEY,
    wallet_address TEXT NOT NULL,
    ticket_id bigint NOT NULL,
    pot_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL default now(),
    updated_at timestamp with time zone NOT NULL default now()
);

-- +migrate StatementBegin

CREATE OR REPLACE FUNCTION update_updated_at_column() RETURNS trigger AS $$
    BEGIN
	    NEW.updated_at = now();
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

-- +migrate StatementEnd

CREATE TRIGGER update_hotpot_user_updated_at_column
    BEFORE UPDATE ON hotpot_user
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();

-- +migrate Down

DROP TABLE IF EXISTS hotpot_user;