DROP TABLE IF EXISTS pokemons;
DROP TRIGGER IF EXISTS validate_types_before_insert;
DROP TRIGGER IF EXISTS validate_types_before_update;

-------------------------------------------------------------------------------

CREATE TABLE pokemons (
  id          integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  name        text    NOT NULL,
  description text    NOT NULL DEFAULT '',
  category    text    NOT NULL DEFAULT '',
  types       text    NOT NULL DEFAULT '',
  abilities   text    NOT NULL DEFAULT ''
);

CREATE TRIGGER validate_types_before_insert
BEFORE INSERT ON pokemons
BEGIN
  SELECT 
    CASE
      WHEN EXISTS (
        SELECT value
        FROM json_each(NEW.types)
        WHERE value NOT IN (
          'BUG', 'DRAGON', 'FAIRY', 'FIRE', 'GHOST', 'GROUND', 'NORMAL', 
          'PSYCHIC', 'STEEL', 'DARK', 'ELECTRIC', 'FIGHTING', 'FLYING', 
          'GRASS', 'ICE', 'POISON', 'ROCK', 'WATER'
        )
      )
      THEN RAISE(ABORT, 'Invalid type in types array')
    END;
END;

CREATE TRIGGER validate_types_before_update
BEFORE UPDATE ON pokemons
BEGIN
  SELECT 
    CASE
      WHEN EXISTS (
        SELECT value
        FROM json_each(NEW.types)
        WHERE value NOT IN (
          'BUG', 'DRAGON', 'FAIRY', 'FIRE', 'GHOST', 'GROUND', 'NORMAL', 
          'PSYCHIC', 'STEEL', 'DARK', 'ELECTRIC', 'FIGHTING', 'FLYING', 
          'GRASS', 'ICE', 'POISON', 'ROCK', 'WATER'
        )
      )
      THEN RAISE(ABORT, 'Invalid type in types array')
    END;
END;
