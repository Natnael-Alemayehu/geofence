-- Version: 1.01
-- Description: Create table geolocation with PostGIS support

-- Enable PostGIS extension (if not already enabled)
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE geolocation (
    location_id TEXT NOT NULL,
    geojson     GEOMETRY NOT NULL,          -- Use PostGIS geometry type
    PRIMARY KEY (location_id)
);