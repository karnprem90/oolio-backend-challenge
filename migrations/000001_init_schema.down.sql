-- Drop indexes
DROP INDEX IF EXISTS idx_products_category;
DROP INDEX IF EXISTS idx_promo_codes_code;

-- Drop tables
DROP TABLE IF EXISTS promo_codes;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products; 