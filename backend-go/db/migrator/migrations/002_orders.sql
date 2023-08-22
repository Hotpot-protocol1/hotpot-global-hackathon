-- +migrate Up

CREATE TABLE IF NOT EXISTS hotpot_order
(
    order_hash TEXT PRIMARY KEY,
    created_at timestamp with time zone NOT NULL default now(),
    updated_at timestamp with time zone NOT NULL default now(),
    sig TEXT NOT NULL,
    chain int NOT NULL default 0,
    offerer TEXT NOT NULL,
    offer_token TEXT NOT NULL,
    offer_token_id numeric(78,0) NOT NULL,
    offer_amount numeric(78,0) NOT NULL,
    end_time numeric(78,0),
    royalty_percent numeric(78,0) NOT NULL DEFAULT 0,
    royalty_recipient TEXT NOT NULL,
    salt numeric(78,0) NOT NULL
);

CREATE TRIGGER update_order_updated_at_column
    BEFORE UPDATE ON hotpot_order
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();

-- +migrate Down

DROP TABLE IF EXISTS hotpot_order;