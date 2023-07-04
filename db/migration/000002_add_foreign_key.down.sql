ALTER TABLE `product` DROP FOREIGN KEY `fk_course_category_id`;
ALTER TABLE `cart`    DROP FOREIGN KEY `fk_cart_user_id`;
ALTER TABLE `cart`    DROP FOREIGN KEY `fk_cart_product_id`;
ALTER TABLE `order`   DROP FOREIGN KEY `fk_order_user_id`;
ALTER TABLE `order`   DROP FOREIGN KEY `fk_order_cart_id`;
