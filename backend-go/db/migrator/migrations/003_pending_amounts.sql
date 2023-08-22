-- +migrate Up

ALTER TABLE user_ticket DROP CONSTRAINT user_ticket_ticket_id_pot_id_key;
ALTER TABLE user_ticket ADD COLUMN pending_amount numeric(78,0);
ALTER TABLE user_ticket ADD CONSTRAINT user_ticket_wallet_address_ticket_id_pot_id_key UNIQUE (wallet_address, ticket_id, pot_id);
CREATE UNIQUE INDEX user_ticket_ticket_id_pot_id_key ON user_ticket (pot_id, ticket_id) WHERE (ticket_id != 0);

-- +migrate Down

DROP INDEX user_ticket_ticket_id_pot_id_key;
ALTER TABLE user_ticket DROP CONSTRAINT user_ticket_wallet_address_ticket_id_pot_id_key;
ALTER TABLE user_ticket DROP COLUMN pending_amount;
ALTER TABLE user_ticket ADD CONSTRAINT user_ticket_ticket_id_pot_id_key UNIQUE (ticket_id, pot_id);