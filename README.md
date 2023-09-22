# Go Fiber MVC App 

This is a Go web application built using the Fiber framework, following the MVC (Model-View-Controller) architectural pattern. The project structure is inspired by Laravel, providing a familiar and organized approach to building web applications.
This Project is fully API based meading currently there is no view implementation here
## Features

- MVC pattern: Organize your code into Models, Views, and Controllers for a structured and maintainable application.
- Go Fiber: Utilize the fast and efficient Fiber web framework for building high-performance web applications.
- Service Layer: Implement a Service layer to encapsulate business logic and keep your controllers lean.
- Flexible Routing: Define routes and handle HTTP requests easily with the powerful routing capabilities of Go Fiber.
- Database Integration: Connect to your preferred database system using GORM, an excellent ORM library for Go.

## Project Structure

The project follows a folder structure similar to Laravel, providing a clear separation of concerns and promoting modular development. Here's a brief overview:

- `app/`: Contains the application-specific code, including models, views, controllers, and services.
- `config/`: Stores configuration files for the application, such as database settings, middleware configurations, etc.
- `database/`:Not Available now but can be customized.
- `public/`: Contains publicly accessible assets such as CSS, JavaScript, and static files.
- `routes/`: Defines the application routes and maps them to the appropriate controllers and actions.
- `main.go`: The entry point of the application where the server starts and routes are registered.

## Getting Started

To get started with the Go Fiber MVC App:

1. Clone the repository.
2. Install the necessary dependencies using `go get` or any package manager you prefer.
3. Configure the application settings in the `config/` folder, such as database connection details.
4. Run database migrations to set up the required database schema.
5. Start the application using `go run main.go` or by building and running the binary.
6. Access the application in your web browser at `http://localhost:3000` or the configured port.

Feel free to explore the code, modify it according to your needs, and build upon this foundation to create powerful web applications using Go Fiber.

