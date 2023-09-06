CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  "name" VARCHAR(255)
);

-- Заполнение таблицы users
INSERT INTO users ("name")
VALUES ('John Doe'), ('Jane Smith'), ('Mike Johnson');

