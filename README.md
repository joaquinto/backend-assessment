# Backend Assessment
This is the Go backend engineering assessment for evolve credit. 
___

## Note:
To run this project locally, make sure you have the following tools installed; Go, docker and prisma. Start up your docker desktop.

After clonning this project, copy the content of what is in the ```.env.sample``` file into your already created ```.env``` file in the root directory.

Run the following command in your terminal sequentially;

```bash
joaquinto@jonathan backend-assessment% docker-compose up -d
joaquinto@jonathan backend-assessment% go mod tidy
joaquinto@jonathan backend-assessment% make prisma-generate
joaquinto@jonathan backend-assessment% make prisma-sync
joaquinto@jonathan backend-assessment% make prisma-seed
joaquinto@jonathan backend-assessment% go run cmd/main.go
 ```

**Features Implemented**
1. Get user
2. Get users (pagination, query by date range)

## API Information
The API endpoints are hosted on Heroku - [Evolve Credit Backend Assessment](https://github.com/joaquinto/backend-assessment)

|METHOD  |DESCRIPTION                        |ENDPOINT                                  |
|------- |-----------------------------------|------------------------------------------|
|GET    |Get User                            |/api/v1/users/:email                       |
|GET    |Get Users (pagination, query by date range)                           |/api/v1/users                        |

|DESCRIPTION         |FIELDS                                                    |                 
|--------------------|-------------------------------------------------------------------|
|Get User             |email (required, parameter)                              |
|Get Users             |limit, pageNumber, from, to (optional, query string)                                                    |


## Author
### Odjegba Jonathan (Joaquinto)