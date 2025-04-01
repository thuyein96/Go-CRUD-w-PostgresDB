# Go CRUD with PostgreSQL

This repository contains a simple CRUD (Create, Read, Update, Delete) API built with Go and PostgreSQL. The implementation follows the tutorial from the YouTube video: [Go CRUD with PostgreSQL](https://www.youtube.com/watch?v=lf_kiH_NPvM&list=WL&index=9&t=705s).

## Features
- RESTful API for managing records
- PostgreSQL as the database
- Clean and structured code with Go
- Uses `gorilla/mux` for routing
- `gorm` as the ORM for database interaction

## Prerequisites
Ensure you have the following installed:
- [Go](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/your-username/Go-CRUD-w-PostgresDB.git
   cd Go-CRUD-w-PostgresDB
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Configure the database:
   - Create a PostgreSQL database
   - Update `.env` file with your database credentials

4. Run the application:
   ```sh
   go run main.go
   ```

## API Endpoints

| Method | Endpoint      | Description        |
|--------|--------------|--------------------|
| GET    | `/items`     | Get all items     |
| GET    | `/items/{id}` | Get item by ID    |
| POST   | `/items`     | Create new item   |
| PUT    | `/items/{id}` | Update item by ID |
| DELETE | `/items/{id}` | Delete item by ID |

## Acknowledgments
- Inspired by [this YouTube tutorial](https://www.youtube.com/watch?v=lf_kiH_NPvM&list=WL&index=9&t=705s)
- Built using Go and PostgreSQL
- Special thanks to Robby for clear and insightful tutorial, which made this implementation possible.


