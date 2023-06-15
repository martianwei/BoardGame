-- 建立 game 表格
CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY,
    player1_id UUID NOT NULL,
    player2_id UUID NOT NULL,
    player3_id UUID NOT NULL,
    player4_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (player1_id) REFERENCES users(id),
    FOREIGN KEY (player2_id) REFERENCES users(id),
    FOREIGN KEY (player3_id) REFERENCES users(id),
    FOREIGN KEY (player4_id) REFERENCES users(id)
);
-- 建立 game_history 表格
CREATE TABLE IF NOT EXISTS game_histories (
    id SERIAL PRIMARY KEY,
    game_id UUID NOT NULL,
    step INT NOT NULL,
    user_id UUID NOT NULL,
    move_from  VARCHAR(255) NOT NULL,
    move_to  VARCHAR(255) NOT NULL,
    FOREIGN KEY (game_id) REFERENCES game(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
