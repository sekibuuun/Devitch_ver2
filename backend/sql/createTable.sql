CREATE TABLE IF NOT EXISTS Streams (
  stream_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  comments TEXT,
  comments_length INTEGER,
  listeners INTEGER,
  listeners_length INTEGER,
  FOREIGN KEY (listeners) REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS Genres (
  genre_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  genre VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Stream_Genres (
  stream_id INTEGER UNSIGNED,
  genre_id INTEGER UNSIGNED,
  PRIMARY KEY (stream_id, genre_id),
  FOREIGN KEY (stream_id) REFERENCES Streams(stream_id),
  FOREIGN KEY (genre_id) REFERENCES Genres(genre_id)
);

CREATE TABLE IF NOT EXISTS Users (
  user_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(100) NOT NULL,
  links JSON
);