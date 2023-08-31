CREATE TABLE  segments (
    id serial PRIMARY KEY,
    slug varchar(255) NOT NULL UNIQUE,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users_segments (
     user_id INT,
     segment_slug varchar(255),
     PRIMARY KEY (user_id, segment_slug),
     created_at timestamptz NOT NULL DEFAULT now()
);
SELECT CURRENT_TIMESTAMP;