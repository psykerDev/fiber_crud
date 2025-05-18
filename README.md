# Fiber CRUD

A simple CRUD (Create, Read, Update, Delete) application built using the Go programming language and the Fiber web framework. This project serves as a foundational example for building RESTful APIs with Go.

## Project Structure

fiber_crud/
├── initializers/ # Initialization logic (DB, configs)
├── models/ # Data models (structs)
├── routes/ # API route handlers
├── tools/ # Utility/helper functions
├── go.mod # Go module dependencies
├── go.sum # Go checksum file
└── server.go # Main application entry point

bash
Copy
Edit

## Getting Started

### Prerequisites

- Go 1.18 or higher
- Git

### Installation

```bash
git clone https://github.com/psykerDev/fiber_crud.git
cd fiber_crud
go mod tidy
go run server.go
The server will start on http://localhost:3000.

Usage
Use tools like Postman or cURL to interact with the API:

Method	Endpoint	Description
GET	/api/items	Retrieve all items
GET	/api/items/:id	Retrieve an item by ID
POST	/api/items	Create a new item
PUT	/api/items/:id	Update an existing item
DELETE	/api/items/:id	Delete an item by ID

Contributing
Contributions are welcome! Please fork the repository, make your changes, and submit a pull request.
