Commands to run the app

1. Build the catalog app
````
docker build -t catalog-app .  
````
2. Run the docker container
````
 docker run -p 8082:8082 -e PORT=8082 catalog-app
````