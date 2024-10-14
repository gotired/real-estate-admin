CREATE TABLE IF NOT EXISTS property_collection (
    id UUID PRIMARY KEY,
    owner UUID NOT NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    type VARCHAR(50),
    location VARCHAR(255),
    price INTEGER NOT NULL,
    photo VARCHAR(255),
    FOREIGN KEY (owner) REFERENCES user_collection(id) ON DELETE CASCADE
);
