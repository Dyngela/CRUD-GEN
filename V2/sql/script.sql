create table product (
    `product_id` INT NOT NULL AUTO_INCREMENT,
    `name` varchar(40) NOT NULL UNIQUE,
    `price` FLOAT(5, 2),
    `since` DATETIME,
    PRIMARY KEY (`product_id`),
);

create table users (
    `user_id` INT NOT NULL AUTO_INCREMENT,
    `name` varchar,
    `product_id` INT NOT NULL,
    PRIMARY KEY (`user_id`),
    CONSTRAINT `Fk_product_id_users` FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`) ON UPDATE CASCADE,
);

-- drop table product;
-- alter table product ADD COLUMN ()
