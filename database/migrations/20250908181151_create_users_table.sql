-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email TEXT NOT NULL UNIQUE,
  username TEXT NOT NULL UNIQUE,
  password VARCHAR(60) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_by UUID REFERENCES users(user_id),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_by UUID REFERENCES users(user_id)
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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS users_set_timestamp_trigger ON users;

DROP FUNCTION IF EXISTS set_timestamp_function();

DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd
