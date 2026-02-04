CREATE TABLE IF NOT EXISTS tokens (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    hash VARCHAR(255) NOT NULL UNIQUE,
    token_type VARCHAR(50) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_tokens_hash ON tokens(hash);
CREATE INDEX idx_tokens_user_id_type ON tokens(user_id, token_type);
CREATE INDEX idx_tokens_expires_at ON tokens(expires_at);