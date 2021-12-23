# locationsvc
Location Service In Golang Using Google Places API

## How to run this service at docker local
1. Run DOCKER DEAMON at your machine successfully
2. Make sure don't have any image mysql is running at port 3306, otherwise you will have 1 error
3. Git clone this repo
4. Comment all values for localhost like
```
    #DB_DRIVER=mysql
    #DB_USER=root
    #DB_PASSWORD=1234567890
    #DB_PORT=3306
    #DB_HOST=127.0.0.1 # For running the app without docker
    #DB_NAME=location_db
```
5. UNcomment all values for docker
```
    #for docker
    DB_DRIVER=mysql
    DB_USER=location_user
    DB_PASSWORD=location_password
    DB_PORT=3306
    DB_HOST=location-mysql
    DB_NAME=location_db
```
5.  Go to terminal at root of project
```sh
   chmod 755 start.sh
   ./start.sh
```

6. If have some logs at console like, server started and worked successfully

```sh
location_db_mysql | Version: '5.7.36'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server (GPL)
location_app      | We are getting the env values
location_app      | We are connected to the mysql database
location_app      | 2021/12/22 01:38:54 /app/src/cmd/db/db.go:28
location_app      | [0.203ms] [rows:-] SELECT DATABASE()
location_app      | 
location_app      | 2021/12/22 01:38:54 /app/src/cmd/db/db.go:28
location_app      | [1.809ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'location_db' AND table_name = 'places' AND table_type = 'BASE TABLE'
location_app      | 
location_app      | 2021/12/22 01:38:54 /go/pkg/mod/gorm.io/driver/mysql@v1.0.5/migrator.go:194
location_app      | [0.283ms] [rows:-] SELECT DATABASE()
location_app      | 
location_app      | 2021/12/22 01:38:54 /go/pkg/mod/gorm.io/driver/mysql@v1.0.5/migrator.go:203
location_app      | [4.346ms] [rows:-] SELECT column_name, is_nullable, data_type, character_maximum_length, numeric_precision, numeric_scale , datetime_precision FROM information_schema.columns WHERE table_schema = 'location_db' AND table_name = 'places'
location_app      | 2021/12/22 01:38:54 Starting server at port:  8091


## How to run this service at localhost
1. Start your mysql at your localhost machine successfully
2. Git clone this repo
3. **Change values in file .env which are mapped with your config (DB_USER, DB_PASSWORD, DB_HOST, SERVER_PORT)**
4. **Change DB_HOST value to 127.0.0.1**
5. Note we set Auth_usename, Auth_password at env file to authenticate APIs, you can check values,
   but remember using it to auth APIs before use
6. **Create database location_db by yourself**
7.  Go to terminal at root of project
```sh
   go get .    
   go run src/cmd/main.go
```
8. If you change something related to swagger, run this command to update documentation,
   it will re-generate docs folder at project source code
```sh
   swag init -g ./src/cmd/main.go
   
   It will say something like 
    2021/12/22 08:38:14 Skipping 'entity.Meta', already parsed.
    2021/12/22 08:38:14 Generating entity.Location
    2021/12/22 08:38:14 Skipping 'entity.Location', already parsed.
    2021/12/22 08:38:14 Generating entity.Pagination
    2021/12/22 08:38:14 Skipping 'entity.Pagination', already parsed.
    2021/12/22 08:38:14 create docs.go at docs/docs.go
    2021/12/22 08:38:14 create swagger.json at docs/swagger.json
    2021/12/22 08:38:14 create swagger.yaml at docs/swagger.yaml

```
9. FE can access this page to see API documentation
   http://localhost:8091/swagger/index.html#/NearByLocations/get_v1_locations_nearby
   ![API documentation swagger](https://i.im.ge/2021/12/22/oe7Ir0.png)



```

## How to test this service
### Testing by run integration tests
At root folder of project, run all ingtegration tests
**You must using localhost to run test**
Uncomment all values for localhost at .env file
```
go test ./src/handler/rest
Return like below means pass all tests

 go test ./src/handler/rest 
ok      github.com/trongtb88/locationsvc/src/handler/rest       3.633s

```
### Testing using swagger
1. FE can access this page to see API documentation
   http://localhost:8091/swagger/index.html#/NearByLocations/get_v1_locations_nearby
   ![API documentation swagger](https://i.im.ge/2021/12/22/oe7Ir0.png)

2. Call APIs at this page, and if return 200, that means API worked fine.

### Testing using Postman
Import
```
curl --location --request GET 'localhost:8091/v1/locations/nearby?street_name=Sukhumvit, Thailand&place_type=restaurant&radius=2&page_token=' \
--header 'Cookie: _csrf=XO_87ETIH7F3gLUPD12HERF3'
```
3. Request success
```
{
    "metadata": {
        "path": "/v1/accounts",
        "status_code": 200,
        "status": "Created",
        "error": {
            "code": "OK",
            "message": "Success"
        },
        "timestamp": "2021-22-12T02:33:44Z"
    },
    "data": [
        {
        "name": "",
        "address" : "" 
        }
    ]
}

```
## Improve
#### Develop API to Authorize when click button Authorize at Swagger
#### using JWT with expired token
#### Build more real authentication flow
#### Build dynamic sql before put to gorm or sql library
#### Logger
#### Add request_id or correlation_id to trace request
#### Add some telemetry to monitors
#### Build CI Pipeline to deploy using Jenkin


