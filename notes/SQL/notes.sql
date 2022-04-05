START TRANSACTION;

CREATE TABLE notes (
  id serial primary key,
  title varchar(100) NOT NULL,
  content text  NOT NULL,
  created date NOT NULL,
  expires date NOT NULL
  --CONSTRAINT idx_snippets_created FOREIGN KEY (created)
  );

INSERT INTO notes (id, title, content, created, expires) VALUES
(1, 'Не имей сто рублей', 'Не имей сто рублей,\nа имей сто друзей.', '2021-01-27 13:09:34', '2022-01-27 13:09:34'),
(2, 'Лучше один раз увидеть', 'Лучше один раз увидеть,\nчем сто раз услышать.', '2021-01-27 13:09:40', '2022-01-27 13:09:40'),
(3, 'Не откладывай на завтра', 'Не откладывай на завтра,\nчто можешь сделать сегодня.', '2021-01-27 13:09:44', '2021-02-03 13:09:44');

COMMIT;
