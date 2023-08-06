-- +migrate Up

CREATE TABLE IF NOT EXISTS user_ticket
(
    id SERIAL PRIMARY KEY,
    wallet_address TEXT NOT NULL,
    ticket_id bigint NOT NULL,
    pot_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL default now(),
    updated_at timestamp with time zone NOT NULL default now(),
    is_winner boolean NOT NULL DEFAULT false,
    raffle_timestamp timestamp with time zone,
    UNIQUE(ticket_id, pot_id)
);

-- +migrate StatementBegin

CREATE OR REPLACE FUNCTION update_updated_at_column() RETURNS trigger AS $$
    BEGIN
	    NEW.updated_at = now();
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

-- +migrate StatementEnd

CREATE TRIGGER update_user_ticket_updated_at_column
    BEFORE UPDATE ON user_ticket
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();

-- +migrate Down

DROP TABLE IF EXISTS user_ticket;