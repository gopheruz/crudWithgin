# Building a RESTful API with Golang and Gin-Gonic

This tutorial demonstrates how to build a simple RESTful API using Golang and the Gin-Gonic web framework.

## Prerequisites

- Go installed on your machine ([Download Go](https://golang.org/dl/))
- Basic understanding of Go programming language

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/nurmuhammaddeveloper/crudWithgin
    cd crudWithgin
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    go mod vendor
    ```

## Usage
#### 1)Add the .env file and fill it with the data as shown in the sample.env file
#### 2)Start the postgres database on your computer
#### 3)After starting postgres create database for project
#### 4)And run migratinos
```bash
    make migrateup
```
#### for delete migration
```bash 
    make migratedown
```
## Run the application:

```bash
go run main.go
```

## How to request?
##### for create user 
```json
{
 "first_name":"John",
 "last_name:":"Doe",
 "email":"johndoe@example.com"
}
```
##### this returns created user as json fromat
```json
{ "id":"UUID",
 "first_name":"John",
 "last_name:":"Doe",
 "email":"johndoe@example.com"
}
```
##### for get user send 
```request
    https://hostname:port/getbyemail/{user_email}
```

##### for delete request type delete
```request
    https://hostname:port/delete/{user_ID}
```
##### for update user request type PUT
```json
{
 "id":"UUID",
 "first_name":"John",
 "last_name:":"Doe",
 "email":"johndoe@example.com"
}
```
##### this returns updated user as json fromat
```json
{ "id":"UUID",
 "first_name":"newName",
 "las_tname:":"newLastname",
 "email":"newname@example.com"
}
```
##### for get many users request type GET
```request
    https://hostname:port/dgetall
```
```json
{
 "page": 1,
 "limit": 10,
 "search":"your search keyword"
}
```
---
##### It returns many user from database using your request parametrs


## Skills Acquired

- **Working with PostgreSQL Database in Golang**: I developed a strong understanding of interacting with PostgreSQL databases using Golang. This included CRUD operations, querying, and managing data.

- **Utilizing Interfaces and Structs**: Leveraging interfaces and structs in Golang allowed me to create efficient and flexible code structures, enabling better data handling and abstraction.

- **Struct-to-JSON Binding**: I learned how to bind Golang structs to JSON representations, facilitating seamless communication between my API and clients.

- **API Development in Golang using Gin-Gonic**: Utilizing the Gin-Gonic framework, I successfully built and deployed robust RESTful APIs, enabling smooth communication between the server and clients.

- **Enhanced Technical Skills**: This project significantly boosted my technical proficiency, allowing me to tackle complex challenges in Golang and database management effectively.

## Technologies Used

1. **GIN-Gonic Library**: Employed Gin-Gonic, a powerful web framework, to create and deploy efficient APIs in Golang.

2. **Database Interaction**: Leveraged the "database/sql" package in conjunction with "github.com/lib/pq" to seamlessly work with PostgreSQL databases, enabling effective data manipulation.

3. **UUID Library Integration**: Integrated the UUID library to generate unique identifiers for user IDs, ensuring uniqueness and reliability in user data management.

