CREATE TABLE IF NOT EXISTS users(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username VARCHAR NOT NULL,
    password_hash VARCHAR NOT NULL
);

INSERT INTO users(username, password_hash) VALUES('Pavlo232', 'e10adc3949ba59abbe56e057f20f883e');

CREATE TABLE IF NOT EXISTS images(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    image_path VARCHAR NOT NULL,
    image_url VARCHAR NOT NULL
);