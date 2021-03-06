# Achievers server

[![Build Status](https://travis-ci.org/ivzb/achievers_server.svg?branch=master)](https://travis-ci.org/ivzb/achievers_server)
[![Go Report Card](https://goreportcard.com/badge/github.com/ivzb/achievers_server)](https://goreportcard.com/report/github.com/ivzb/achievers_server)
[![GoDoc](https://godoc.org/github.com/ivzb/achievers_server?status.svg)](https://godoc.org/github.com/ivzb/achievers_server) 

Achievers Web API in Go

To download, run the following command:

~~~
go get github.com/ivzb/achievers_server
~~~

## Quick Start with PostgreSQL 

Please ensure all necessary plugins are available by following:

```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

Create database via executing sql script in postgre
It is located in /config/postgre.sql
Enter postgre and execute

```
\i {path_to_sql_script}
```

Open config/config.json and edit the Database section so the connection information matches your Postgre instance.

Build and run from the root directory. Open your REST client to: http://localhost. You should see the welcome message and status 200.

## Quick Start with Docker

Build and run the image

Invoke Docker from the achievers_server package directory to build an image using the Dockerfile: 

```
$ docker build -t achievers_server .
```

This will fetch the golang base image from Docker Hub, copy the package source to it, build the package inside it, and tag the resulting image as achievers_server.

To run a container from the resulting image:

```
$ docker run --publish 8080:8080 --name achievers_server --rm achievers_server
```

The --publish flag tells docker to publish the container's port 8080 on the external port 8080.

The --name flag gives our container a predictable name to make it easier to work with.

The --rm flag tells docker to remove the container image when the outyet server exits. 

With the container running, open http://localhost:8080/ in a web browser and you should see welcome message.

Now that we've verified that the image works, shut down the running container from another terminal window:

```
$ docker stop achievers_server 
```

## Available Endpoints

The following endpoints are available:

```
* GET  /v1/ - Retrieve an welcome message

* POST /v1/user/create - Create new user
* POST /v1/user/auth   - Retrieve an access token

* GET  /v1/profile?id= - Retrieve profile by ID
* GET  /v1/profile/me  - Retrieve your own profile

* POST /v1/achievement/create                 - Create new achievement
* GET  /v1/achievement?id=                    - Retrieve an achievement by ID
* GET  /v1/achievements/last                  - Retrieve last page of achievements
* GET  /v1/achievements/after?after_id=       - Retrieve page of achievements after specified one
* GET  /v1/achievements/quest?id=&page=       - Retrieve achievements by quest_id and page
* GET  /v1/achievements/quest/last            - Retrieve last page of achievements by quest_id
* GET  /v1/achievements/quest/after?after_id= - Retrieve page of achievements by quest_id after specified one

* POST /v1/evidence/create           - Create new evidence 
* GET  /v1/evidence?id=              - Retrieve an evidence by ID
* GET  /v1/evidences/last            - Retrieve last page of evidences
* GET  /v1/evidences/after?after_id= - Retrieve page of evidences after specified one

* POST /v1/reward/create           - Create new reward 
* GET  /v1/reward?id=              - Retrieve a reward by ID
* GET  /v1/rewards/last            - Retrieve last page of rewards
* GET  /v1/rewards/after?after_id= - Retrieve page of rewards after specified one

* POST /v1/quest/create           - Create new quest 
* GET  /v1/quest?id=              - Retrieve a quest by ID
* GET  /v1/quests?page=           - Retrieve quests by page
* GET  /v1/quests/last            - Retrieve last page of rewards
* GET  /v1/quests/after?after_id= - Retrieve page of rewards after specified one

* POST /v1/quest_achievement/create - Create a new quest_achievement

* POST /v1/file/create - Create new file
* GET  /v1/file?id=    - Retrieve a file by ID
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

## TODO // bigger number == higher priority
1. add kubernetes support
2. implement constructTest for app/shared/ similar to this in controller/
3. implement router with following format: router.GET("path", handler, middleware)
4. use router's handlers to build a help page with endpoints
5. improve context - replace model instance with model type
6. improve code coverage
7. consider unioning db_test's exists and existsMultiple functions (invent smart exister interface)
8. implement microservices architecture

## Done :) // some tasks might not be written
0. quest_achievement controller
+    1. create quest_achievement
+    2. get achievements by quest_id
+    3. get quests by achievement_id
1. logger - add log error method
2. get achievements by quest_id
3. extract common controller functionalities which return plain Result
 handlers should have \*shared.Request instead of \*http.Request which will wrap \*http.Request and use it only internally
4. extract consts from controller/controller.go to shared/consts.go
5. improve paging concept (pass afterID)
6. extract framework models from app models
7. replace paging concept with later/after as implemented in achievement controller
8. update readme with available endpoints
9. refactor app/db/ as it should recieve some abstraction
10. extract limit const from db.go to config and pass it when instantiate
11. refactor afterID to id in controllers
12. refactor db's context - extract needed info from model's properties
13. get rid of db's scan function, use struct tags instead
14. improve db's create - extract needed fields by struct's tags
15. create form validation
16. unify db length constraints
