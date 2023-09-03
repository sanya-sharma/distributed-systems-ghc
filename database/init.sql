
-- Create the "users" table to store user information
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username TEXT NOT NULL,
  password TEXT NOT NULL,
  role TEXT NOT NULL
);

-- Create the "customers" table to store customer information
CREATE TABLE IF NOT EXISTS customers (
   id SERIAL PRIMARY KEY,
   name TEXT NOT NULL,
   email TEXT NOT NULL,
   phone_number TEXT NOT NULL,
   address TEXT NOT NULL,
   user_id INTEGER NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP,
   FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Create the "products" table to store product information
CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  price NUMERIC(10, 2) NOT NULL,
  stock_qty INT NOT NULL
    );

-- Create the "orders" table to store order information
CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL,
  order_date TIMESTAMP NOT NULL,
  status TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE
    );

-- Create the "order_items" table to store individual items within an order
CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id)
    );

-- Create the "payments" table to store payment details for each order
CREATE TABLE IF NOT EXISTS payments (
 id SERIAL PRIMARY KEY,
 order_id INT NOT NULL,
 payment_method TEXT NOT NULL,
 amount NUMERIC(10, 2) NOT NULL,
 status TEXT NOT NULL,
 FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE
    );
