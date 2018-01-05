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
* POST   /user/auth       - Retrieve an access token
* POST   /user/create      - Create a new user

* GET	 /achievement?id= - Retrieve an achievement by ID
* GET	 /achievements?page= - Retrieve achievements by page
* GET    /achievements/quest?id=&page= - Retrieve achievements by quest_id and page
* POST   /achievement/create - Create a new achievement

* GET	 /evidence?id= - Retrieve an evidence by ID
* GET	 /evidences?page= - Retrieve evidences by page
* POST   /evidence/create - Create a new evidence 

* GET	 /reward?id= - Retrieve a reward by ID
* GET	 /rewards?page= - Retrieve rewards by page
* POST   /reward/create - Create a new reward 

* GET	 /quest?id= - Retrieve a quest by ID
* GET	 /quests?page= - Retrieve quests by page
* POST   /quest/create - Create a new quest 

* POST   /quest_achievement/create - Create a new quest_achievement
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
go test ./... -cover
```

Run specific package tests with coverage

```
go test ./app/controller -coverprofile=coverage.out
```

View coverage result in html

```
go tool cover -html=coverage.out
```

## DB

Create database via executing sql script in mysql
It is localted in /config/mysql.sql
Enter mysql and execute

```
source {path_to_sql_script}
```

## TODO
1. extract common controller functionalities which return plain Result

