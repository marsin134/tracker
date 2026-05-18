CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS routes
(
      route_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
      route_speed REAL DEFAULT 0,
      route_average_speed REAL DEFAULT 0,
      route_way REAL DEFAULT 0
);

CREATE INDEX idx_route_user_id ON routes(user_id);