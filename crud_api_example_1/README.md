## Go REST API with Gin and MongoDB

A minimal CRUD API built with Go, Gin, and MongoDB. It exposes basic user operations and demonstrates clean project structure with controllers, models, and database package.

### Features

- Gin-based HTTP server with grouped routes
- MongoDB integration using the official Go driver
- Unique index on `user.id`
- Simple, composable controllers

### Tech Stack

- Go (Gin, MongoDB Go Driver)
- MongoDB

### Project Structure

```
crud_api_example_1
  main.go
  database/
    database.go
  models/
    userModel.go
  controllers/
    user/
      GetUsers.go
      InsertUser.go
      DeleteUser.go
      UpdateUser.go
```

### Prerequisites

- Go 1.20+
- MongoDB instance (local or remote)

### Environment Variables

Create a `.env` file in the project root:

```
MONGODB_URI=mongodb://localhost:27017
```

On startup, the app connects to MongoDB, selects database `go_test_db`, and ensures a unique index on `user.id`.

### Install & Run

```
go mod tidy
go run main.go
```

Production tip:

```
export GIN_MODE=release
go run main.go
```

### API Reference

Base URL: `http://localhost:8080/api`

#### Get users (limited)

- GET `/users/:size`
- Params: `size` (path, integer) — max number of users to return
- Response: 200 JSON array of users

```
curl http://localhost:8080/api/users/10

# Sample response
[
  {
    "id": 1,
    "isActive": true,
    "balance": "$1,234.56",
    "picture": "https://example.com/pic.jpg",
    "age": 30,
    "name": "John Doe",
    "gender": "male",
    "company": "ACME",
    "email": "john@example.com",
    "phone": "+1 111 222 3333",
    "address": "1 Main St",
    "about": "Sample user",
    "registered": 1704067200000,
    "latitude": 41.0,
    "longitude": 29.0,
    "favoriteFruit": "apple"
  }
]
```

#### Create user

- POST `/users/add`
- Body (JSON): see User model below
- Response: 200 insert result (MongoDB)

```
curl -X POST http://localhost:8080/api/users/add \
  -H 'Content-Type: application/json' \
  -d '{
    "id": 1,
    "isActive": true,
    "balance": "$1,234.56",
    "picture": "https://example.com/pic.jpg",
    "age": 30,
    "name": "John Doe",
    "gender": "male",
    "company": "ACME",
    "email": "john@example.com",
    "phone": "+1 111 222 3333",
    "address": "1 Main St",
    "about": "Sample user",
    "registered": 1704067200000,
    "latitude": 41.0,
    "longitude": 29.0,
    "favoriteFruit": "apple"
  }'

# Sample response
{ "InsertedID": "<mongo_object_id>" }
```

Postman sample body (raw, JSON):

```json
{
  "id": 1001,
  "isActive": true,
  "balance": "$2,450.00",
  "picture": "https://picsum.photos/seed/gonzalo/200/200",
  "age": 29,
  "name": "Gonzalo Gortinez",
  "gender": "male",
  "company": "Gortinez Labs",
  "email": "gonzalo.gortinez@example.com",
  "phone": "+54 9 11 5555 5555",
  "address": "Av. Siempre Viva 742, Buenos Aires",
  "about": "Early adopter and backend engineer.",
  "registered": 1704067200000,
  "latitude": -34.6037,
  "longitude": -58.3816,
  "favoriteFruit": "banana"
}
```

#### Update user (partial)

- PUT `/users/:id/update`
- Params: `id` (path, integer)
- Body: partial JSON (fields to update). `id` and `_id` are ignored by design.
- Response: 200 updated user document

```
curl -X PUT http://localhost:8080/api/users/1/update \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "new@example.com",
    "isActive": false,
    "favoriteFruit": "banana"
  }'

# Sample response
{
  "id": 1,
  "isActive": false,
  "balance": "$1,234.56",
  "picture": "https://example.com/pic.jpg",
  "age": 30,
  "name": "John Doe",
  "gender": "male",
  "company": "ACME",
  "email": "new@example.com",
  "phone": "+1 111 222 3333",
  "address": "1 Main St",
  "about": "Sample user",
  "registered": 1704067200000,
  "latitude": 41.0,
  "longitude": 29.0,
  "favoriteFruit": "banana"
}
```

#### Delete user

- DELETE `/users/:id/delete`
- Params: `id` (path, integer)
- Response: 200 confirmation JSON

```
curl -X DELETE http://localhost:8080/api/users/1/delete

# Sample response
{ "message": "User deleted successfully" }
```

### Data Model

`models/userModel.go`

### Indexes

On startup, the app creates a unique index on `user.id`:

```
db.user.createIndex({ id: 1 }, { unique: true })
```

Duplicate inserts will return a duplicate key error from MongoDB.

### Notes

- `.env` is required; the app will exit if `MONGODB_URI` is missing
- For local development, add MONGODB_URI=mongodb://localhost:27017/ to your .env file.
- All user operations target the `user` collection in `go_crud_db`
- Change database/collection names in `database/database.go` if needed
