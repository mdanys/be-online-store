ALTER TABLE `product` ADD CONSTRAINT `fk_product_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);
ALTER TABLE `cart`    ADD CONSTRAINT `fk_cart_user_id`        FOREIGN KEY (`user_id`)     REFERENCES `user`     (`id`);
ALTER TABLE `cart`    ADD CONSTRAINT `fk_cart_product_id`     FOREIGN KEY (`product_id`)  REFERENCES `product`  (`id`);
ALTER TABLE `order`   ADD CONSTRAINT `fk_order_user_id`       FOREIGN KEY (`user_id`)     REFERENCES `user`     (`id`);
ALTER TABLE `order`   ADD CONSTRAINT `fk_order_cart_id`       FOREIGN KEY (`cart_id`)     REFERENCES `cart`     (`id`);
