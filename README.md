Hello

## Part 1: ibrary

Before running the application, run the following command to install and start the necessary Docker containers:

sh

    docker compose up

This will install and start two Docker containers: postgres-byfood and pgadmin-byfood. Make sure these containers are running before launching the application.

Run migrate/main.go to create the models in the database

### Dependencies

The application uses the following packages:

* Gin: "github.com/gin-gonic/gin" for routing
* GORM: for interacting with the database
* Validator: "github.com/go-playground/validator/v10" for validation
* Log: "log" for logging into log/log.log

##  Part 2: url

#### /url 
This endpoint accepts JSON input and returns results as specified. The accepted JSON format is:
        
        {
            "url": "https://BYFOOD.com/food-EXPeriences?query=abc/",
            "operation": "canonical"
        }
#### Response
The endpoint returns a response in the following format:
       
        {
            "result": {
                    "Processed_url": "https://www.byfood.com/food-experiences"
                }
        }

The application then stores the operation in the database table urls.

    create table urls
    (
    url_id      bigserial
    primary key,
    url_given   text not null,
    operation   text not null,
    url_treated text not null
    );
