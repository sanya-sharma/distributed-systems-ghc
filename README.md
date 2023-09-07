# distributed-systems-ghc
This is the master repo for GHC23 presentation


**Docker commands to start the database and order app from the Docker Compose:**

1. Build the order docker app
````
cd order
go build -o main . 
````
2. Run the docker compose command
````
cd ..
docker-compose up --build
````
