CREATE TABLE IF NOT EXISTS categories(
    id UUID PRIMARY KEY,
    name VARCHAR(64),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS posts(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    title VARCHAR(100),
    body TEXT,
    category_id UUID REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS comments(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    post_id UUID REFERENCES posts(id),
    body TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS tags(
    id UUID PRIMARY KEY,
    name VARCHAR(64),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS post_tags(
    id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts(id),
    tag_id UUID REFERENCES tags(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);