-- enum type for roles
CREATE TYPE role_enum AS ENUM ('admin', 'client');


-- uuid does not auto generate here

-- enum type for genres
CREATE TYPE genre_enum AS ENUM ('Action', 'Comedy', 'Drama', 'Horror', 'Sci-Fi', 'Romance', 'Thriller');
ALTER TYPE genre_enum ADD VALUE 'Fantasy';

-- users table
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    pass VARCHAR(255) NOT NULL,
    role role_enum NOT NULL
);

-- movies table
CREATE TABLE movies (
    id VARCHAR(36) PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,
    creator VARCHAR(100) NOT NULL,
    genre genre_enum NOT NULL,
    release_date DATE NOT NULL,
    desc_movie TEXT NULL,
    poster_image_url TEXT NULL, 
    ticket_price DECIMAL(10, 2) NOT NULL
);

-- showtimes
CREATE TABLE showtimes (
    id VARCHAR(36) PRIMARY KEY,
    movie_id VARCHAR(36) NOT NULL,
    add_by VARCHAR(36) NOT NULL,
    date_start TIMESTAMP NOT NULL,
    date_end TIMESTAMP NOT NULL,
    CONSTRAINT fk_showtimes_movie FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
    CONSTRAINT fk_showtimes_user FOREIGN KEY (add_by) REFERENCES users(id) ON DELETE CASCADE
);

-- seats
CREATE TABLE seats (
    id VARCHAR(36) PRIMARY KEY,
    showtime_id VARCHAR(36) NOT NULL,
    seat VARCHAR(20) NOT NULL,
    status VARCHAR(50) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    CONSTRAINT fk_seats_showtime FOREIGN KEY (showtime_id) REFERENCES showtimes(id) ON DELETE CASCADE,
    CONSTRAINT fk_seats_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

ALTER Table users ADD constraint unique_user UNIQUE (username)

ALTER Table movies DROP COLUMN genre

ALTER Table movies ADD COLUMN genre genre_enum[]