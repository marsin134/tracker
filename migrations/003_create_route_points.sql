CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS route_points (
    id BIGSERIAL PRIMARY KEY,
    route_id UUID NOT NULL REFERENCES routes(route_id) ON DELETE CASCADE,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_route_points_route_id ON route_points(route_id);
CREATE INDEX idx_route_points_route_id_created_at ON route_points(route_id, created_at);