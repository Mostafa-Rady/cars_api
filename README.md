# Cars-API

> Simple RESTfUl API to manage cars persistence. 

## Start app

Requires docker and docker-compose, to edit db schema "db/migrations/init.sql"

``` bash
# requires docker and docker-compose
$ sudo docker-compose up

```
## Run tests
``` bash
# Test
$ go test ./... -coverprofile cover.out

# view coverage report
$ go tool cover -html=cover.out
```

``` bash
go build
./go_restapi
```

## Endpoints
import included postman collection "cars.postman_collection.json"
