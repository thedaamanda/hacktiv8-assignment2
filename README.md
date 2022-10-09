# Hacktiv8 Golang Assignment 2 - Scalable Web Service with Golang

This project is a REST API created using Golang Programming Language with Gin Gonic Framework, GORM for object relational mapping for PostgreSQL database and Swaggo for API documentation

## Setup Config

Add the keys you created to your `.env` file. You can copy the [`.env.example`](.env.example) file and fill in the blanks.

```
APP_HOST=xxx
APP_PORT=xxx

DB_HOST=xxx
DB_USER=xxx
DB_PASSWORD=xxx
DB_DATABASE=xxx
DB_SSLMODE=xxx
```

## Available Endpoints

|  Method | URL | Description |
| ------------ | ------------ | ------------ |
| GET | /orders  | Get all `Orders` from database |
| POST | /orders  | Create an `Order` to database |
| GET | /orders/:id | Get an `Order` by order_id |
| PUT | /orders/:id | Update an `Order`  |
| DELETE | /orders/:id | Delete an `Order` from database |
| GET | /swagger/*any | Swagger/OpenAPI 2.0 documentations |
| GET | /swagger/index.html | Swagger/OpenAPI 2.0 documentations index page |
