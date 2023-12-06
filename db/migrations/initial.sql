CREATE TABLE IF NOT EXISTS Lists (
  list_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  qtd_users smallint NOT NULL DEFAULT (0) CHECK(qtd_users >= 0 AND qtd_users <= 100)
);

CREATE TABLE IF NOT EXISTS Users (
  user_id VARCHAR(255) PRIMARY KEY,
  username VARCHAR(255) NOT NULL CHECK(username <> ''),
  email VARCHAR(255) NOT NULL CHECK(email <> ''),
  user_hash VARCHAR(255) NOT NULL CHECK(user_hash <> ''),
  celphone VARCHAR(255) NOT NULL CHECK(celphone <> ''),
  user_status BOOLEAN NOT NULL,
  list_id UUID REFERENCES Lists (list_id)
);

CREATE OR REPLACE FUNCTION addUser(
    id VARCHAR(255),
    username VARCHAR(255),
    email VARCHAR(255),
    user_hash VARCHAR(255),
    celphone VARCHAR(255),
    user_status BOOLEAN
) RETURNS UUID AS $$
DECLARE
    list_record Lists;
BEGIN
    
	SELECT * INTO list_record FROM Lists WHERE qtd_users < 100 LIMIT 1;
  IF list_record IS NULL THEN
    INSERT INTO Lists (qtd_users) VALUES (0) RETURNING * INTO list_record;
	END IF;

  INSERT INTO Users (user_id, username, email, user_hash, celphone, user_status, list_id)
    VALUES (id, username, email, user_hash, celphone, user_status, list_record.list_id);

  UPDATE Lists SET qtd_users = list_record.qtd_users + 1 WHERE list_id = list_record.list_id;

  RETURN list_record.list_id;
END;
$$ LANGUAGE plpgsql;