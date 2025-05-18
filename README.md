📦 Project Name
Fiber CRUD

📝 Description
A simple CRUD (Create, Read, Update, Delete) REST API built with Go and the Fiber web framework. A starter template for learning and developing APIs with Go.

🗂️ Project Structure
python
Copy
Edit
fiber_crud/
├── initializers/       # Initialization logic (DB, configs)
├── models/             # Data models (structs)
├── routes/             # API route handlers
├── tools/              # Utility/helper functions
├── go.mod              # Go module dependencies
├── go.sum              # Go checksum file
└── server.go           # Main application entry point
🚀 Getting Started
Prerequisites
Go 1.18+

Git

Installation
bash
Copy
Edit
git clone https://github.com/psykerDev/fiber_crud.git
cd fiber_crud
go mod tidy
go run server.go
Server starts at: http://localhost:3000

🛠️ Usage
HTTP Method	Endpoint	Description
GET	/api/items	Retrieve all items
GET	/api/items/:id	Retrieve item by ID
POST	/api/items	Create a new item
PUT	/api/items/:id	Update an existing item
DELETE	/api/items/:id	Delete an item

🤝 Contributing
Feel free to fork, make changes, and open a pull request!

📄 License
MIT License

