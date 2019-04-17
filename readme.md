# Address Book Assignment

## How to run the application

1. Clone the application with `git@github.com:krishsb2405/assignment.git`

2. Use the mySQL dump `sql-dump.sql` to create the database, create the tables and insert some dump data.

3. Once the application is cloned and database is created, change the Connection String as per your mysql username, password and database name on line No. 20 in api.go.

4. Download the postman configuration file (optional).

5. There are number of dependencies which need to be imported before running the application. Please get the dependenices through the following commands -

    ```shell
        go get "github.com/go-sql-driver/mysql"
        go get "github.com/gorilla/mux"
    ```

6. To run the application, please use the following command -

    ```shell
        go run api.go entry.go
    ```
> Note: By default the port number its being run on is **12345**.

## Endpoints Description

### Get All Entries

```
    URL - *http://localhost:12345/api/entries*
    Method - GET
```

### Get Entry By ID

```JSON
    URL - *http://localhost:12345/api/entry?id=1*
    Method - GET
```

### Create Entry

```JSON
    URL - *http://localhost:12345/api/entry*
    Method - POST
    Body - (content-type = application/json)
    {
    	"first_name":"John",
    	"last_name":"Doe",
    	"email_address":"john.doe@gmail.com",
    	"phone_number":"1234567890"
    }
```

### Update Entry

```JSON
    URL - *http://localhost:12345/api/entry*
    Method - PUT
    Body - (content-type = application/json)
    {
        "id":5,
    	"first_name":"John",
    	"last_name":"Doe",
    	"email_address":"john.doe@gmail.com"
    }
```

### Delete Entry

```JSON
    URL - *http://localhost:12345/api/entry?id=1*
    Method - DELETE
```

### Import entries through CSV

```JSON
    URL - *http://localhost:12345/api/upload-entries-CSV*
    Method - POST
    Body - form-data
    {
        "csvFile": <csv-file>
    }
```

### Download entries to CSV

```JSON
    URL - *http://localhost:12345/api/download-entries-CSV*
    Method - GET
```

## Test Driven Development Description

To run all the unit test cases, please do the following -

1. `go run api.go entry.go`
2. `go test -v`


## Hope everything works. Thank you.