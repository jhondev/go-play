CREATE TABLE authors (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  bio text
);
CREATE TABLE games (
  id BIGSERIAL PRIMARY KEY,
  key_name text NOT NULL,
  name text NOT NULL
);
CREATE TABLE matches (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  score text
);