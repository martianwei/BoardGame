-- CREATE TYPE military_rank_enum AS ENUM (
--   '軍旗' = 0,
--   '司令',
--   '軍長',
--   '師長',
--   '旅長',
--   '團長',
--   '營長',
--   '連長',
--   '排長',
--   '工兵'
-- );



CREATE TABLE IF NOT EXISTS formations (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    strategy int[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);