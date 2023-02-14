-- +goose Up

CREATE TABLE IF NOT EXISTS users(
        id BIGINT NOT NULL AUTO_INCREMENT,
        client_id VARCHAR(191) NOT NULL,
        client_secret LONGTEXT NOT NULL,
        grant_type LONGTEXT NOT NULL,
        access_token LONGTEXT,
        refresh_token LONGTEXT,
        redirect_uri LONGTEXT NOT NULL,
        created_at DATETIME(3),
        updated_at DATETIME(3),
        deleted_at DATETIME(3),
        PRIMARY KEY(id)
);



CREATE TABLE IF NOT EXISTS unisender_users(
      id BIGINT AUTO_INCREMENT PRIMARY KEY,
      account_id BIGINT UNIQUE NOT NULL,
      token longtext NOT NULL,
      uname longtext NOT NULL,
      created_at DATETIME(3),
      updated_at DATETIME(3),
      deleted_at DATETIME(3)
);

CREATE TABLE IF NOT EXISTS contacts(
       id BIGINT AUTO_INCREMENT PRIMARY KEY,
       contact_id BIGINT NOT NULL,
       account_id BIGINT NOT NULL,
       name longtext NOT NULL,
       email VARCHAR(255) UNIQUE,
       created_at DATETIME(3),
       updated_at DATETIME(3),
       deleted_at DATETIME(3),
       FOREIGN KEY (account_id)
            REFERENCES unisender_users(account_id)
                    ON UPDATE CASCADE
                    ON DELETE CASCADE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd


-- +goose Down
DROP TABLE IF EXISTS contacts;

DROP TABLE IF EXISTS unisender_users;

DROP TABLE IF EXISTS users;
-- +goose StatementBegin
SELECT 'down SQL query';


-- +goose StatementEnd
