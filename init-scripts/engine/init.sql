CREATE TABLE engines (
  id SERIAL PRIMARY KEY,
  car_id INTEGER UNIQUE,
  "type" VARCHAR(255),
  horsepower INTEGER
);

-- Заполнение таблицы engines
INSERT INTO engines (car_id, "type", horsepower)
VALUES
  (1, 'V6', 300),
  (2, 'Inline-4', 250),
  (3, 'V8', 450);
