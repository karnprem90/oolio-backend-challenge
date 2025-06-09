-- Remove sample orders
DELETE FROM orders WHERE id IN ('0000-0000-0000-0001', '0000-0000-0000-0002');

-- Remove sample promo codes
DELETE FROM promo_codes WHERE code IN ('HAPPYHRS', 'FIFTYOFF', 'WELCOME20', 'SUMMER25', 'WINTER30');

-- Remove sample products
DELETE FROM products WHERE id IN ('1', '2', '3', '4', '5', '6', '7', '8', '9', '10'); 