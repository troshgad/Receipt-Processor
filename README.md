# Files of Note:

- processor.go
  - handles all the routes / processing of the requests
- validation.go
  - validates if the request body is a valid receipt 
#  Prerequisites:

- Have Docker installed on device
- Have Postman
- Clone Repository

if you do not have Postman or Docker, you can install them here:
 -  <a>https://www.docker.com/</a>
 - <a>https://www.postman.com/</a>
## Build image for Docker:

- Navigate with the terminal to the project directory that was cloned previously
- Once there run this command to build image

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
1. http://localhost:8080/receipts/process
```

```
2. http://localhost:8080/receipts/{id}/points
```
- Endpoint 1 is a POST request that takes a JSON payload

an example of the format is below:
```
  {
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
      {
        "shortDescription": "Mountain Dew 12PK",
        "price": "6.49"
      },{
        "shortDescription": "Emils Cheese Pizza",
        "price": "12.25"
      },{
        "shortDescription": "Knorr Creamy Chicken",
        "price": "1.26"
      },{
        "shortDescription": "Doritos Nacho Cheese",
        "price": "3.35"
      },{
        "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
        "price": "12.00"
      }
    ],
    "total": "35.35"
  }
```

you will get an id in the response of the first endpoint and that id will be used to fulfill the id parameter of the second endpoint. an example of the response of the first endpoint is shown below:

```
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```
An example of the request in postman is shown below:
![image](https://github.com/user-attachments/assets/c5b2c9ea-a6b6-44c1-a438-53ffe585f81d)

- Endpoint 2 is a GET request that takes an id in the URI and will return back JSON with the number of points that the receipt is awarded. An example of the request URI is shown below:

```
http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points
```
The returned value will look like this:
```
{ "points": 32 }
```

