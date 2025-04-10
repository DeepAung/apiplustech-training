-- tables
-- Table: movies
CREATE TABLE movies (
    id integer NOT NULL CONSTRAINT movies_pk PRIMARY KEY AUTOINCREMENT,
    title varchar(255) NOT NULL
);

-- Table: reviews
CREATE TABLE reviews (
    id integer NOT NULL CONSTRAINT reviews_pk PRIMARY KEY,
    stars int NOT NULL,
    comment varchar(255) NOT NULL,
    movie_id int NOT NULL,
    CONSTRAINT reviews_movies FOREIGN KEY (movie_id)
    REFERENCES movies (id)
);

-- End of file.
