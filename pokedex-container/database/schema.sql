DROP TABLE IF EXISTS pokemons;
DROP FUNCTION IF EXISTS validate_pokemon_types;
DROP TRIGGER IF EXISTS validate_types_before_insert ON pokemons;
DROP TRIGGER IF EXISTS validate_types_before_update ON pokemons;

-------------------------------------------------------------------------------

CREATE TABLE pokemons (
  id          SERIAL  NOT NULL,
  name        TEXT    NOT NULL UNIQUE,
  description TEXT    NOT NULL DEFAULT '',
  category    TEXT    NOT NULL DEFAULT '',
  types       JSONB   NOT NULL DEFAULT '[]',
  abilities   JSONB   NOT NULL DEFAULT '[]'
);

CREATE OR REPLACE FUNCTION validate_pokemon_types()
RETURNS TRIGGER AS $$
DECLARE
  type_value TEXT;
BEGIN
  FOR type_value IN SELECT jsonb_array_elements_text(NEW.types)
  LOOP
    IF type_value NOT IN (
      'BUG', 'DRAGON', 'FAIRY', 'FIRE', 'GHOST', 'GROUND', 'NORMAL',
      'PSYCHIC', 'STEEL', 'DARK', 'ELECTRIC', 'FIGHTING', 'FLYING',
      'GRASS', 'ICE', 'POISON', 'ROCK', 'WATER'
    ) THEN
      RAISE EXCEPTION 'Invalid type in types array: %', type_value;
    END IF;
  END LOOP;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER validate_types_before_insert
BEFORE INSERT ON pokemons
FOR EACH ROW
EXECUTE FUNCTION validate_pokemon_types();

CREATE TRIGGER validate_types_before_update
BEFORE UPDATE ON pokemons
FOR EACH ROW
EXECUTE FUNCTION validate_pokemon_types();
