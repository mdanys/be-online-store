CREATE TABLE `user` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `email`                 VARCHAR(255)    NOT NULL UNIQUE,
    `password`              VARCHAR(255)    NOT NULL,
    `name`                  VARCHAR(255)    NOT NULL,
    `role`                  VARCHAR(255)    NOT NULL,
    `dob`                   DATE,
    `gender`                VARCHAR(1),
    `address`               TEXT,
    `user_picture`          VARCHAR(1024),
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `category` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `name`                  VARCHAR(255)    NOT NULL,
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `product` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `category_id`           BIGINT          NOT NULL,
    `name`                  VARCHAR(255)    NOT NULL,
    `price`                 DECIMAL(10,2)   NOT NULL,
    `qty`                   BIGINT          NOT NULL,
    `rating`                DECIMAL(2,1),
    `detail`                TEXT,
    `product_picture`       VARCHAR(1024),
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);


CREATE TABLE `cart` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `user_id`               BIGINT          NOT NULL,
    `product_id`            BIGINT          NOT NULL,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `order` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `cart_id`               BIGINT          NOT NULL,
    `total_price`           DECIMAL(10,2)   NOT NULL,
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
