-- 建立 game 表格
CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 建立 game_history 表格
CREATE TABLE IF NOT EXISTS game_histories (
    id SERIAL PRIMARY KEY,
    game_id UUID NOT NULL,
    step INT NOT NULL,
    move_from  VARCHAR(255) NOT NULL,
    move_to  VARCHAR(255) NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games(id)
);

-- 建立 game_user 表格
CREATE TABLE IF NOT EXISTS game_formations (
    id SERIAL PRIMARY KEY,
    game_id UUID NOT NULL,
    formation_id UUID NOT NULL,
    move_order INT NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games(id),
    FOREIGN KEY (formation_id) REFERENCES formations(id)
);
