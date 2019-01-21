# Address Book Assignment

## How to run the application

1. Clone the application with git@github.com:krishsb2405/assignment.git

2. Use the mySQL dump to create the database, create the tables and insert some dump data.

3. Once the application is cloned and database is created, change the Connection String on line No. 20 in api.go.

4. There are number of dependencies which need to be imported before running the application. Please get the dependenices through the following commands -

    ```shell
        go get "github.com/go-sql-driver/mysql"
        go get "github.com/gorilla/mux"
    ```

5. To run the application, please use the following command -

    ```shell
        go run api.go entry.go
    ```
