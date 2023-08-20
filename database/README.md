# distributed-systems-ghc Database setup
This is the master repo for GHC23 presentation


**Docker commands to start the database from the volume:**

1. Build the docker image
   docker build -t postgres_db_image .
2. Create the docker volume
   docker volume create postgres_db_data
3. Runt the container
   docker run -d -p 5432:5432 --name postgres_db_container -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=db --mount source=postgres_db_data,target=/var/lib/postgresql/data  postgres_db_image
4. Check if the container is running 
   docker ps -a --filter "name=postgres_db_container"