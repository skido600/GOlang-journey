# Go Fundamentals & Web Development Examples

This repository serves as a comprehensive collection of Go programming examples, delving into core language features and illustrating their application in building a basic web application. It's designed to help understand Go's syntax, concurrency models, data structures, and how to integrate with popular web frameworks and databases.

## Features

*   **Go Language Fundamentals**: Demonstrations covering basic data types (strings, integers, booleans, floats), control structures (`if/else`, `switch`, `for` loops), variable and constant declarations, and `iota`.
*   **Data Structures**: Examples of working with arrays, slices, and maps, including creation, manipulation, and iteration.
*   **Concurrency**: Introduction to Goroutines for lightweight concurrency and Channels for safe communication between concurrent processes.
*   **Pointers & Memory Management**: Explanations and examples of Go's pointer system, including address-of and dereference operators.
*   **Functions & Interfaces**: Examples showcasing function definitions, return values, error handling, and basic interface usage.
*   **Error Handling**: Standard Go error handling patterns using `error` interfaces.
*   **Math & String Utilities**: Usage of Go's built-in `math` and `strings` packages for common operations like square root, power, absolute value, string manipulation, and random number generation.
*   **HTTP Server Development**: A basic HTTP server built using Go's `net/http` package.
*   **RESTful API with Gin**: Implementation of a simple RESTful API using the Gin Web Framework for handling HTTP requests and responses.
*   **Database Integration with GORM**: Demonstrates connecting to a PostgreSQL database and defining models using GORM (Go Object-Relational Mapper).
*   **Environment Variable Management**: Utilization of `godotenv` for loading and managing environment variables, essential for configuration.

## Stacks / Technologies

| Technology       | Description                                  | Link                                                 |
| :--------------- | :------------------------------------------- | :--------------------------------------------------- |
| **Go**           | Programming Language                         | [https://golang.org/](https://golang.org/)           |
| **Gin Gonic**    | High-performance HTTP web framework          | [https://gin-gonic.com/](https://gin-gonic.com/)     |
| **GORM**         | The fantastic ORM library for Go             | [https://gorm.io/](https://gorm.io/)                 |
| **PostgreSQL**   | Powerful, open source object-relational DB   | [https://www.postgresql.org/](https://www.postgresql.org/) |
| **Go DotEnv**    | Load environment variables from `.env` file  | [https://github.com/joho/godotenv](https://github.com/joho/godotenv) |

## Installation

To get a local copy up and running, follow these simple steps.

### Prerequisites

*   Go (version 1.18 or higher recommended)
    *   [Install Go](https://golang.org/doc/install)
*   PostgreSQL (for the `crud` module)
    *   [Install PostgreSQL](https://www.postgresql.org/download/)

### Setup

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
    cd YOUR_REPO_NAME
    ```
    (Replace `YOUR_USERNAME` and `YOUR_REPO_NAME` with actual values if this were a real GitHub repo).

2.  **Navigate to specific modules**:
    The repository is structured into different modules/directories. You'll need to navigate into them to run their respective `go.mod` files.

3.  **Install Go dependencies**:
    For the `crud` module:
    ```bash
    cd crud
    go mod tidy
    ```
    For the `http` module:
    ```bash
    cd http
    go mod tidy
    ```
    And for any other module that has a `go.mod` file.

4.  **Set up Environment Variables**:
    For modules like `crud` and `http` that use `godotenv`, you'll need to create a `.env` file in their respective root directories.

    **`crud/.env` example**:
    ```
    DB_URL="host=localhost user=postgres password=12345 dbname=demopost port=5432 sslmode=disable"
    ```
    Adjust the `DB_URL` with your PostgreSQL connection details.

    **`http/.env` example**:
    ```
    # No specific variables are strictly required for the basic HTTP example,
    # but the setup is there if you extend it.
    ```

## Usage

Each directory represents a different concept or mini-project.

### Running Individual Go Examples

To run any of the fundamental Go examples (e.g., `Goroutine`, `aboutArraysINGo`, `control`, `dataType`, `deep_pointer`, `errorhandlin`, `math`, `structure_main`, `variable && constant`, `hello_world`, `string && method`), navigate into its directory and execute the `main.go` file:

```bash
cd Goroutine
go run main.go
# Output: lorem202k
```

### Running the `crud` Web Application

1.  Ensure your PostgreSQL database is running and accessible with the credentials provided in `crud/.env`.
2.  Navigate to the `crud` directory:
    ```bash
    cd crud
    ```
3.  Run the application:
    ```bash
    go run main.go
    ```
4.  The server will start on `http://localhost:8080`. You can test the default endpoint:
    ```bash
    curl http://localhost:8080/
    # Expected output: {"message":"pong"}
    ```
    (Further CRUD endpoints would need to be implemented in `crud/main.go` and `crud/controllers/` etc.)

### Running the Basic `http` Server

1.  Navigate to the `http` directory:
    ```bash
    cd http
    ```
2.  Run the application:
    ```bash
    go run main.go
    ```
3.  The server will start on `http://localhost:8080`. You can test the default endpoint:
    ```bash
    curl http://localhost:8080/
    # Expected output: Hello World
    ```

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

[![Readme was generated by Readmit](https://img.shields.io/badge/Readme%20was%20generated%20by-Readmit-brightred)](https://readmit.vercel.app)
