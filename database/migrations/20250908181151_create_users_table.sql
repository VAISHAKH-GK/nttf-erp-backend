-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL UNIQUE,
  password VARCHAR(60) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_by UUID REFERENCES users(id)
);

CREATE TABLE account_types (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_by UUID REFERENCES users(id),
  is_active BOOLEAN
);

CREATE TABLE user_account_types (
  user_id UUID REFERENCES users(id),
  account_type_id UUID REFERENCES account_types(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_by UUID REFERENCES users(id),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_by UUID REFERENCES users(id),
  PRIMARY KEY(user_id, account_type_id)
);

CREATE OR REPLACE FUNCTION set_timestamp_function() RETURNS TRIGGER AS $$
BEGIN
  IF TG_OP = 'INSERT' THEN
    NEW.created_at = NOW();
    NEW.updated_at = NOW();
  ELSIF TG_OP = 'UPDATE' THEN
    NEW.updated_at = NOW();
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER users_set_timestamp_trigger BEFORE
INSERT OR UPDATE ON users FOR EACH ROW EXECUTE FUNCTION set_timestamp_function();

CREATE TRIGGER account_types_set_timestamp_trigger BEFORE
INSERT OR UPDATE ON account_types FOR EACH ROW EXECUTE FUNCTION set_timestamp_function();

CREATE TRIGGER user_account_types_set_timestamp_trigger BEFORE
INSERT OR UPDATE ON user_account_types FOR EACH ROW EXECUTE FUNCTION set_timestamp_function();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS user_account_types_set_timestamp_trigger ON user_account_types;
DROP TRIGGER IF EXISTS account_types_set_timestamp_trigger ON account_types;
DROP TRIGGER IF EXISTS users_set_timestamp_trigger ON users;

DROP FUNCTION IF EXISTS set_timestamp_function();

DROP TABLE IF EXISTS user_account_types CASCADE;
DROP TABLE IF EXISTS account_types CASCADE;
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd
