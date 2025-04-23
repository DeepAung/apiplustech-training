CREATE TABLE pokemons (
  id          BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name        text    NOT NULL,
  description text    NOT NULL DEFAULT '',
  category    text    NOT NULL DEFAULT '',
  types       text    NOT NULL DEFAULT '',
  abilities   text    NOT NULL DEFAULT ''
);
