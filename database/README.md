# distributed-systems-ghc Database setup
This is the master repo for GHC23 presentation


**Docker commands to start the database from the volume:**

1. Build the docker image
````
docker build -t postgres_db_image .
````
2. Create the docker volume
````
docker volume create postgres_db_data
````
3. Runt the container
````
docker run -d -p 5432:5432 --name postgres_db_container -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=db --mount source=postgres_db_data,target=/var/lib/postgresql/data  postgres_db_image
   ````
4. Check if the container is running 
````
docker ps -a --filter "name=postgres_db_container"
   ````
5. Run the following command to enter the Postgres CLI
````
docker exec -it postgres_db_container psql -U user -d db
   ````
6. Run the following sql to insert data into users table
````
INSERT INTO users (username, password, role)
   VALUES
   ('user1', 'password1', 'user'),
   ('user2', 'password2', 'user'),
   ('user3', 'password2', 'user');
 ````
7. Run the following sql to insert data into customer table
````
   INSERT INTO customers (name, email, phone_number, address, user_id, created_at, updated_at)
   VALUES
   ('John Doe', 'johndoe@example.com', '123-456-7890', '123 Main St, Anytown, USA', 1, CURRENT_TIMESTAMP, NULL),
   ('Jane Smith', 'janesmith@example.com', '987-654-3210', '456 Elm St, Othertown, USA', 2, CURRENT_TIMESTAMP, NULL),
   ('Bob Johnson', 'bjohnson@example.com', '555-123-4567', '789 Oak St, Somewhere, USA', 3, CURRENT_TIMESTAMP, NULL);
````