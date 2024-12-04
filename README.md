#  Prerequisites:

- Have Docker installed on device
- Clone Repository

## Build image for Docker:

- Go to project directory that was cloned previously 
- Run command to build image

```
  docker build -t docker-gs-ping:multistage -f Dockerfile.multistage .
```

## Run Docker Container: 

- After building the image, use this commmand to run the image

```
docker run --publish  8080:8080  docker-gs-ping:multistage
```

You should now be able to reach the app located at </a> http://localhost:8080 </a>

## NOTES:

The two endpoints are

```
http://localhost:8080/receipts/process
```

and 

```
http://localhost:8080/receipts/{id}/points
```

you will get an id in the response of the first endpoint and that id will be used to fulfill the id parameter of the second endpoint.

To test the endpoints, I would use Postman linked here : 
  <a>https://www.postman.com/</a>
