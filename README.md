# distributed-systems-ghc
This is the master repo for GHC23 presentation


**Docker commands to start the database from the volume:**

1. docker volume create my_postgres_data
2. docker run -d -p 5432:5432 --name my_postgres_container \
   -e POSTGRES_USER=myuser \
   -e POSTGRES_PASSWORD=mypassword \
   -e POSTGRES_DB=mydb \
   --mount source=my_postgres_data,target=/var/lib/postgresql/data \
   my_postgres_image