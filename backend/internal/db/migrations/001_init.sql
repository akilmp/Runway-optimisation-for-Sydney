CREATE TABLE weather_obs (
  id SERIAL PRIMARY KEY,
  metar TEXT NOT NULL,
  taf TEXT NOT NULL,
  observed_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE notams (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  valid_from TIMESTAMP,
  valid_to TIMESTAMP
);
