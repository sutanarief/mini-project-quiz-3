-- +migrate Up
-- +StatementBegin

CREATE TABLE category (
  id SERIAL PRIMARY KEY,
  name VARCHAR(256),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE book (
  id SERIAL PRIMARY KEY,
  title VARCHAR(256),
  description VARCHAR(500),
  image_url VARCHAR(500),
  release_year INT,
  price VARCHAR(50),
  total_page INT,
  thickness VARCHAR(20),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  category_id INT
);

-- +migrate StatementEnd