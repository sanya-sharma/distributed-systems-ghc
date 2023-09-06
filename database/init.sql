
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

-- Create the "catalogs" table to store product information
CREATE TABLE IF NOT EXISTS catalogs (
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
    FOREIGN KEY (product_id) REFERENCES catalogs (id)
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

INSERT INTO users (username, password, role)
VALUES
    ('user1', 'password1', 'user'),
    ('user2', 'password2', 'user'),
    ('user3', 'password2', 'user');

INSERT INTO customers (name, email, phone_number, address, user_id, created_at, updated_at)
VALUES
    ('John Doe', 'johndoe@example.com', '123-456-7890', '123 Main St, Anytown, USA', 1, CURRENT_TIMESTAMP, NULL),
    ('Jane Smith', 'janesmith@example.com', '987-654-3210', '456 Elm St, Othertown, USA', 2, CURRENT_TIMESTAMP, NULL),
    ('Bob Johnson', 'bjohnson@example.com', '555-123-4567', '789 Oak St, Somewhere, USA', 3, CURRENT_TIMESTAMP, NULL);

INSERT INTO catalogs (name, description, price, stock_qty) VALUES
    ('Product 4', 'Cotton saree solid', 1999.99, 100),
    ('Product 5', 'Cotton saree printed’', 2999.99, 150),
    ('Product 6', 'Solid chiffon saree’', 5999.99, 200);
