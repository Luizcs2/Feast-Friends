-- TODO - Luiz
-- 1. Create users table with all required fields
-- 2. Create posts table with user foreign key
-- 3. Create follows table for social relationships
-- 4. Create likes table for post interactions
-- 5. Create comments table for post discussions
-- 6. Create events table for social events
-- 7. Create event_rsvps table for event participation
-- 8. Create conversations and messages tables for DMs
-- 9. Add all necessary indexes for performance
-- 10. Set up foreign key constraints


-- 1. Users Table
CREATE Table users {
    id UUID PRIMARY KEY,
    email VARCHAR UNIQUE,
    username VARCHAR UNIQUE,
    full_name VARCHAR,
    bio TEXT,
    profile_picture_url VARCHAR,
    followers_count INT DEFAULT 0,
    following_count INT DEFAULT 0,
    posts_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
}

-- 2. Posts Table
CREATE TABLE posts{
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    title VARCHAR(255),
    image_url VARCHAR,
    description TEXT,
    recipe JSONB, --{ingredients: [], instructions: []}
    likes_count INT DEFAULT 0,
    comments_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
}

CREATE TABLE follows (
    id UUID PRIMARY KEY,
    follower_id UUID REFERENCES users(id),
    following_id UUID REFERENCES users(id),
    created_at TIMESTAMP,
    UNIQUE(follower_id, following_id)
);

CREATE TABLE likes (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    post_id UUID REFERENCES posts(id),
    created_at TIMESTAMP,
    UNIQUE(user_id, post_id)
);

CREATE TABLE comments (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    post_id UUID REFERENCES posts(id),
    content TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Events system
CREATE TABLE events (
    id UUID PRIMARY KEY,
    creator_id UUID REFERENCES users(id),
    title VARCHAR(255),
    description TEXT,
    location VARCHAR,
    event_date TIMESTAMP,
    max_attendees INTEGER,
    current_attendees INTEGER DEFAULT 0,
    image_url VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE event_rsvps (
    id UUID PRIMARY KEY,
    event_id UUID REFERENCES events(id),
    user_id UUID REFERENCES users(id),
    status VARCHAR CHECK (status IN ('attending', 'maybe', 'not_attending')),
    created_at TIMESTAMP,
    UNIQUE(event_id, user_id)
);

-- Direct messaging
CREATE TABLE conversations (
    id UUID PRIMARY KEY,
    participant_1 UUID REFERENCES users(id),
    participant_2 UUID REFERENCES users(id),
    last_message_at TIMESTAMP,
    created_at TIMESTAMP,
    UNIQUE(participant_1, participant_2)
);

CREATE TABLE messages (
    id UUID PRIMARY KEY,
    conversation_id UUID REFERENCES conversations(id),
    sender_id UUID REFERENCES users(id),
    content TEXT,
    message_type VARCHAR DEFAULT 'text',
    read_at TIMESTAMP,
    created_at TIMESTAMP
);