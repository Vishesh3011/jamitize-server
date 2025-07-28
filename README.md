<div align="center">
  <h1>jamitize-server</h1>
</div>

### Introduction
Welcome to the jamitize-server repository! This project is designed to provide a robust backend solution for the Jamitize application, enabling users to create, manage, improve and share their music journey and connect with each other.

### Project Structure
```
├── cmd/
│   ├── api/
│   │   └── main.go
├── data
├── db
└── internal
    ├── core/
    │   ├── application/
    │   ├── config/
    │   ├── middleware/
    ├── controller/
    ├── clients/
    ├── models/
    ├── repository/
    ├── routes/
    ├── service/
    ├── types/
    └── utils/
```

### Prerequisites
MongoDB is required to run this project. Ensure you have it installed and running on your local machine or server.

### Setup authentication for database
Login to Mongodb locally and create the database and user for the application.
```
mongosh
use jamitizedb
```
Create a user in MongoDB
```
db.createUser({
  user: "jamitize-user",
  pwd: "jamitize-user",
  roles: [
    { role: "readWrite", db: "jamitizedb" }
  ]
})
```
Enable authentication in MongoDB by editing the `mongod.conf` file:
```
security:
  authorization: enabled
```
Check connection using:
```
mongosh -u jamitize-user -p jamitize-user --authenticationDatabase jamitizedb
```

### Setup schema and pre-requisites inside database
```bash
mongosh -u jamitize-user -p jamitize-user --authenticationDatabase jamitizedb < db/schema.js

```

### Running the Application




