# Audiience Backend code challenge

This is a little REST microservice created in Go with SQLite as DB, I decided to go the dependency injection way. Uses middlewares for validation and authorization based on valid IP address.


# Run with docker

```
 docker build -t app .
```
running, the app listens on the 8080 port

```
 docker run -it -p 8080:8080 app
```
# Sending request

The service has a single endpoint, '/estimate' it calculates the final estimation and requires the path params "state", "distance", "type" and "base_amount", it also requires the custom header "ip-client" using a valid ip address, here's an example of a request using curl:

```
 curl -X GET "http://localhost:8080/estimate?state=NY&type=premium&distance=22&base_amount=342" -H 'ip-client: 127.0.0.0' 

```
