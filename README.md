# MyApp

MyApp is a sample Go application using the Gin framework, designed with a clean and scalable project structure.

## Table of Contents

- [Project Structure](#project-structure)
- [Requirements](#requirements)
- [Setup](#setup)
- [Running the Application](#running-the-application)
- [Configuration](#configuration)
- [License](#license)

## Project Structure
-   myapp/
    ├── cmd/
    │ └── server/
    │ └── main.go
    ├── config/
    │ └── config.go
    ├── controllers/
    │ └── user_controller.go
    ├── models/
    │ └── user.go
    ├── routes/
    │ └── routes.go
    ├── services/
    │ └── user_service.go
    ├── utils/
    │ └── response.go
    ├── go.mod
    ├── go.sum
    ├── config/
    │ └── config.json

## Requirements

- [Go](https://golang.org/dl/) 1.18 or later
- [Gin](https://github.com/gin-gonic/gin) 1.7.4 or later
- [Viper](https://github.com/spf13/viper) 1.9.0 or later

## Setup

1. **Clone the repository:**

   ```sh
   git clone https://github.com/yourusername/myapp.git
   cd myapp

2. **Install dependencies:**
    ```go mod tidy

3. **Create configuration file:**
    ```
    Ensure you have a config.json file in the config directory with the necessary configuration.

## Running the Application

- To start the application, use the following command:
    go run cmd/server/main.go
    By default, the application will run on port 8080. You can change the port in the config/config.json file.

## Configuration

- The configuration file (config/config.json) should contain the following settings:
    {
        "port": "8081"
    }

## License

- This project is licensed under the MIT License. See the LICENSE file for details.

