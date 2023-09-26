# Building Resilient Distributed Systems With Golang
This repo is designed solely for the purpose of GHC'23 Level Up Lab titled 'Building Resilient Distributed Systems With Golang'.

Presenters:<br />
    1. [Aayushi Chadha](https://www.linkedin.com/in/aayushi-chadha/)<br />
    2. [Sanya Sharma](https://www.linkedin.com/in/sanyasharma2511/)<br />
    3. [Utsha Sinha](https://www.linkedin.com/in/utsha-sinha1510)<br />

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
postman --version
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

In this section, we'll be implementing the circuit breaker and add the relevant code to the payment [service.go](https://github.com/sanya-sharma/distributed-systems-ghc/blob/main/payment/service/service.go) file

**1. Add the structure that implements the circuit breaker**

```
    /* 
       CircuitBreaker implements the circuit breaker
       It has two fields:
        1. mu - Locks, Unlocks the instance of variable.
        2. open - Stores the state of circuit (open/close)
    */
    type CircuitBreaker struct {
        mu   sync.Mutex
        open bool
    }
```

**2. Add the implementation of ExecuteTransaction for executing circuit breaker**<br />
    &nbsp;TODOs:<br />
      &ensp;- If the circuit is open, we will not attempt to call the gateway.<br />
      &ensp;- When encountered failure, if the number of consecutive failures is greater than the desired failure count open the circuit.<br />
      &ensp;- Best Practice: Print log informing user of state of circuit wherever required.<br />

```
    /* 
       ExecuteTransaction executes the transaction using circuit breaker.
       Given a certain number of failures happening in a payment gateway, 
       it'll stop retrying to avoid undue stress on the system and break open the circuit.
       Circuit will be reset after a certain interval of time.
    */
    func (cb *CircuitBreaker) ExecuteTransaction(operation func() bool, consecutiveFails int, paymentGateway string) bool {
        cb.mu.Lock()
        defer cb.mu.Unlock()
        if cb.open {
            // TODO
        }

        completed := operation()

        if !completed {
            // TODO
        }

        return completed
    }
```

**3. Add the  ResetAfterDelay function that closes the circuit after certain time**<br />
    &nbsp;TODOs:<br />
      &ensp;- Make the system take a sleep for sometime.<br />
      &ensp;- Close the circuit so the gateway is ready for another set of requests<br />
      &ensp;- Using goroutine, call the ResetAfterDelay function from ExecuteTransaction while opening the circuit so that it resets automatically after some time.<br />
```
    // ResetAfterDelay resets the circuit after a delay.
    func (cb *CircuitBreaker) ResetAfterDelay(paymentGateway string) {
        // TODO
    }
    
```

**4. Modify the InitiatePayment function to call above circuit breaker code to execute payment**
```
for retry := 0; retry <= maxRetries; retry++ {
    if retry != 0 {
        // Log the retry and sleep before the next attempt
        log.Printf("Payment gateway %v is unavailable. Retrying payment, attempt %d", paymentGateway, retry)
        time.Sleep(time.Second * time.Duration(retry))
    }

    completed = circuit.ExecuteTransaction(func() bool {
        return paymentContext.ExecutePayment()
    }, retry, paymentGateway)
    if completed {
        break
    }
}
```
