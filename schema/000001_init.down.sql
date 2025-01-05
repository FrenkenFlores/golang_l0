-- Drop the items table first as it depends on orders
DROP TABLE IF EXISTS items;

-- Drop the payment table as it depends on orders
DROP TABLE IF EXISTS payment;

-- Drop the delivery table as it depends on orders
DROP TABLE IF EXISTS delivery;

-- Drop the orders table
DROP TABLE IF EXISTS orders;
