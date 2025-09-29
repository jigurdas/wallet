-- migrations/0001_init.down.sql
-- Drop initial schema (reverse of 0001_init.up.sql)

DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS users;
