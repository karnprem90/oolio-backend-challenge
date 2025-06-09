-- Insert sample products
INSERT INTO products (id, name, price, category) VALUES
    ('1', 'Chicken Waffle', 12.99, 'Waffle'),
    ('2', 'Classic Waffle', 8.99, 'Waffle'),
    ('3', 'Belgian Waffle', 10.99, 'Waffle'),
    ('4', 'Chocolate Waffle', 11.99, 'Waffle'),
    ('5', 'Strawberry Waffle', 11.99, 'Waffle'),
    ('6', 'Bacon Waffle', 13.99, 'Waffle'),
    ('7', 'Blueberry Waffle', 12.99, 'Waffle'),
    ('8', 'Banana Waffle', 11.99, 'Waffle'),
    ('9', 'Nutella Waffle', 12.99, 'Waffle'),
    ('10', 'Caramel Waffle', 11.99, 'Waffle');

-- Insert sample promo codes
-- These are example codes that should be valid according to the rules
INSERT INTO promo_codes (code, file_name) VALUES
    ('HAPPYHRS', 'couponbase1.gz'),
    ('HAPPYHRS', 'couponbase2.gz'),
    ('FIFTYOFF', 'couponbase1.gz'),
    ('FIFTYOFF', 'couponbase3.gz'),
    ('WELCOME20', 'couponbase2.gz'),
    ('WELCOME20', 'couponbase3.gz'),
    ('SUMMER25', 'couponbase1.gz'),
    ('SUMMER25', 'couponbase2.gz'),
    ('WINTER30', 'couponbase2.gz'),
    ('WINTER30', 'couponbase3.gz');

-- Insert some sample orders
INSERT INTO orders (id, items, products) VALUES
    ('0000-0000-0000-0001', 
     '[{"productId": "1", "quantity": 2}, {"productId": "3", "quantity": 1}]',
     '[{"id": "1", "name": "Chicken Waffle", "price": 12.99, "category": "Waffle"}, {"id": "3", "name": "Belgian Waffle", "price": 10.99, "category": "Waffle"}]'
    ),
    ('0000-0000-0000-0002',
     '[{"productId": "2", "quantity": 1}, {"productId": "4", "quantity": 2}]',
     '[{"id": "2", "name": "Classic Waffle", "price": 8.99, "category": "Waffle"}, {"id": "4", "name": "Chocolate Waffle", "price": 11.99, "category": "Waffle"}]'
    ); 