CREATE SCHEMA arena;

CREATE TABLE arena.authors (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  bio text
);
CREATE TABLE arena.games (
  id BIGSERIAL PRIMARY KEY,
  key_name text NOT NULL,
  name text NOT NULL
);
CREATE TABLE arena.matches (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  score text
);