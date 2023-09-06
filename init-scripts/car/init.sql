CREATE TABLE cars (
  id SERIAL PRIMARY KEY,
  brand VARCHAR(255),
  model VARCHAR(255),
  "year" INTEGER,
  user_id INTEGER
);

-- Заполнение таблицы cars
INSERT INTO cars (brand, model, "year", user_id)
VALUES
  ('Toyota', 'Cammry', 2021, 1),
  ('Honda', 'Accord', 2022, 2),
  ('Ford', 'Mustang', 2020, 3);
