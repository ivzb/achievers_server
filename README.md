# Achievers server

[![Build Status](https://travis-ci.org/ivzb/achievers_server.svg?branch=master)](https://travis-ci.org/ivzb/achievers_server)
[![Go Report Card](https://goreportcard.com/badge/github.com/ivzb/achievers_server)](https://goreportcard.com/report/github.com/ivzb/achievers_server)
[![GoDoc](https://godoc.org/github.com/ivzb/achievers_server?status.svg)](https://godoc.org/github.com/ivzb/achievers_server) 

Achievers Web API in Go

To download, run the following command:

~~~
go get github.com/ivzb/achievers_server
~~~

If you are on Go 1.5, you need to set GOVENDOREXPERIMENT to 1. If you are on Go 1.4 or earlier, the code will not work because it uses the vendor folder.

## Quick Start with MySQL

Start MySQL and import config/mysql.sql to create the database and tables.

Open config/config.json and edit the Database section so the connection information matches your MySQL instance.

Build and run from the root directory. Open your REST client to: http://localhost. You should see the welcome message and status 200.

To create a user, send a POST request to http://localhost/user with the following fields: first_name, last_name, email, and password.

## Available Endpoints

The following endpoints are available:

```
* POST   http://localhost/auth       - Retrieve an access token

* POST   http://localhost/users		 - Create a new user
* GET	 http://localhost/users/{id} - Retrieve a user by ID
* GET	 http://localhost/users 	 - Retrieve a list of all users
* PUT	 http://localhost/users/{id} - Update a user by ID
* DELETE http://localhost/users/{id} - Delete a user by ID
* DELETE http://localhost/users		 - Delete all users
```

## Rules for Consistency

Rules for mapping HTTP methods to CRUD:

```
POST   - Create (add record into database)
GET    - Read (get record from the database)
PUT    - Update (edit record in the database)
DELETE - Delete (remove record from the database)
```

Rules for status codes:

```
* Read something - 200 (OK)
* Update something - 200 (OK)
* Delete something - 200 (OK)
* Create something - 201 (Created)
* Create but missing info - 400 (Bad Request)
* Access w/ invalid token - 401 (Unauthorized)
* Any other error - 500 (Internal Server Error)
```

Rules for messages:

```
* 200 - item found; no items to find; items deleted; no items to delete; etc
* 201 - item created
* 400 - [field] is missing; [field] needs to be type: [type]
* 401 - unauthorized
* 500 - an error occurred, please try again later (should also log error because it's a programming or server issue)
```

## Tests

Run all tests

```
go test ./vendor/app/...
```

Run specific package tests with coverage

```
go test ./vendor/app/controller -coverprofile=coverage.out
```

View coverage result in html

```
go tool cover -html=coverage.out
```

## DB

Create database via executing sql script in mysql
It is localted in /config/mysql.sql

```
source {path_to_sql_script}
```

## TODO
1. get achievements by quest_id
2. extract common controller functionalities which return plain Result

