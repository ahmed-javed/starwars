# Starwars REST API application

This API can perform CRUD operations for starships.

The entire application can be used by running `main.go` file.

`.env.dist` is a configuration file where application port and MySQL can be configured.

You can use `Makefile` to build and run the entire application.

## Install

    make build

## Run the app

    make run

# REST API Endpoints

The API Endpoints examples are given below.

## Test if the application is listening to the port

### Request

`GET /`

    curl -i -H 'Accept: application/json' http://localhost:9000/

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Tue, 05 Oct 2021 20:30:48 GMT
    Content-Length: 0

## Create a Starcraft

### Request

`POST /spaceship`

    curl --location --request POST 'http://localhost:9000/spaceship' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "Devastator 2",
        "class": "Star Destroyer 2",
        "armament_id": "",
        "crew": 35000,
        "image": "https:\\\\url.to.image",
        "value": 1999.99,
        "status": "operational 2"
    }'

### Response

    {
        "success": true,
        "message": "Spaceship created successfully.",
        "data": null
    }

## Get list of starships

### Request

`GET /spaceships`

    curl --location --request GET 'http://localhost:9000/spaceships'

### Response
    {
        "success": true,
        "message": "",
        "data": [
            {
                "id": 1,
                "name": "Devastator",
                "class": "Star Destroyer",
                "armament_id": "",
                "crew": 35000,
                "image": "https:\\\\url.to.image",
                "value": 1999.99,
                "status": "operational"
            },
            {
                "id": 2,
                "name": "CR90 corvette",
                "class": "Corvette",
                "armament_id": "",
                "crew": 165,
                "image": "https:\\\\url.to.image",
                "value": 3500000,
                "status": "operational"
            },
            .......................
        ]
    }

## Update spaceship

### Request

`PUT /spaceship`

    curl --location --request PUT 'http://localhost:9000/timezones' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImU0N2M5Y2Q3LWE0MjgtNGVjNC1hMmQ3LTFjNDg3ODgyZTk1NyIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYzMzQ2ODk1OCwidXNlcl9pZCI6MTZ9.TDaCSu4j4Jx-U0RUhNLEcmsOGFh-ekJet9RO8uEQa3c' \
        --data-raw '{
            "id": 929,
            "title": "Breakfast",
            "city": "Banyuurip",
            "zone": "Asia/Jakarta",
            "gmt": "GMT+3"
        }'

### Response

#### When the id is found and request is successful
    curl --location --request PUT 'http://localhost:9000/spaceship' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "id":7,
        "name": "Devastator 4",
        "class": "Star Destroyer 2",
        "crew": 35000,
        "image": "https:\\\\url.to.image",
        "value": 1999.99
    }'

#### Provided id not found
    {
        "success": false,
        "error": {
            "code": 4,
            "message": "Unable to update spacecrafts.",
            "exception": "Record not found."
        }
    }

#### Something wrong with the payload
    {
        "success": false,
        "error": {
            "code": 1,
            "message": "Cannot parse request body.",
            "exception": "invalid character '}' looking for beginning of object key string"
        }
    }

## Delete spaceship

### Request

`DELETE /spaceship`

    curl --location --request DELETE 'http://localhost:9000/spaceship' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "id":7
    }'

### Response

#### Spaceship deleted successfully
    {
        "success": true,
        "message": "Spacecraft deleted.",
        "data": null
    }
    
#### Record not found
    {
        "success": true,
        "message": "",
        "data": "Unable to delete spacecrafts."
    }
