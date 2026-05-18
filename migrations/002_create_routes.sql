CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS routes
(
      route_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      route_speed REAL DEFAULT 0,
      route_average_speed REAL DEFAULT 0,
      route_way REAL DEFAULT 0
);