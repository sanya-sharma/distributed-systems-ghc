# distributed-systems-ghc
This is the master repo for GHC23 presentation


### Lab Prerequisites

**1. Check if Go (Golang) is installed:**

Open a terminal and run the following command:

````
go version
````
If Go is installed, you will see output that shows the Go version. If it's not installed, follow these instructions to install it:

* Visit the official Go download page: https://golang.org/dl/
* Download the installer for your operating system.
* Follow the installation instructions for your specific OS.


**2. Check if Docker is installed:**

Run the following command in the terminal:

````
docker --version
````
If Docker is installed, it will display the Docker version. If it's not installed, follow these instructions to install it:

* Visit the official Docker download page: https://docs.docker.com/get-docker/
* Download the Docker Desktop for your operating system.
* Install Docker following the installation instructions provided for your OS.


**3. Check if Postman is installed:**

Run the following command in the terminal:

````
postman --version`
````

If Postman is installed, it will display the Postman version. If it's not installed, follow these instructions to install it:

* Visit the official Postman download page: https://www.postman.com/downloads/
* Download the Postman app for your operating system.
* Install Postman following the installation instructions provided for your OS.
* Use this link to access sample requests for the lab: [Postman GHC Collection](https://winter-star-7764.postman.co/workspace/GHC~d573817e-ed58-47c3-9649-154b689c53a5/collection/29024639-a2ec43b9-7243-4c17-9e70-c146c0b26dab?action=share&creator=29024639)


### Lab Activity 1:
**1. Open your docker desktop app or run the following command**
```
open -a Docker 
```
**2. Run the docker compose command on the terminal**
````
docker-compose up --build
````
**3. Ensure the applications are up and running by looking at the docker logs**

**4. Import the postman collection** 
* Open this link: [Postman GHC Collection](https://winter-star-7764.postman.co/workspace/GHC~d573817e-ed58-47c3-9649-154b689c53a5/collection/29024639-a2ec43b9-7243-4c17-9e70-c146c0b26dab?action=share&creator=29024639)
* Click on the three dots next to the GHC Collection and fork the collection
* Open your Postman app, you should be able to view the collection. (Please ensure you are using the same login on the app and web view)

**5. Test the Get Catalog API request**

View the API response to see all the saree varieties available along with their description, retail price and quantity available

**6. Test the Place Order API request**

Navigate to the [Payment Service](https://github.com/sanya-sharma/distributed-systems-ghc/blob/main/payment/service/service.go#L16). As you can see the service retries each of the available payment gateway available

### Lab Activity 2:

**1. Add changes**