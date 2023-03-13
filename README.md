# Henry Digiexam test
This is the Digiexam test by Henry Cardoso.
I used this website to know how to calculate the [GPA](https://www.scholaro.com/gpa-calculator/)

## Database
The database is in a Docker container to start it run the command:
```shell
docker compose up
```

## Run migrations
Run the migrations using the command in Makefile.
```shell
make migrate
```

## Create migrations file
Create the migrations file using CLI command.
```shell
migrate create -ext sql -dir database/postgres/migration -seq create_table_NAME
```

## Running Tests
To run the tests, the database needs to be up and running and have sure to run the migrations using the above commands.
```shell
go test ./... -v
```

### Endpoints

## GET

# Get 
■ The brand wants to create X number of discount codes
Adding the headers `Blg-Brand-Id` and `Blg-Brand-Name` is necessary. It will check if the brand id and name are hardcoded in the middleware and it is inserted in the migrations.

`http://localhost:9596/brand/discount`

Example of curl command:
```shell
curl -X POST http://localhost:9596/brand/discount -H "Blg-Brand-Id: 1" -H "Blg-Brand-Name: brand1" -H 'Content-Type: application/json' -d '{"amountDiscount":4}'
```

