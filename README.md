#  go-test-mobydev

This is a simple web application - test project - written in Go 

## Prerequisites

Before running this application, make sure you have the following installed on your machine:

- Go programming language (version 1.16 or higher)
- Docker (optional, if you want to run the application in a Docker container)

## Running the Application

### 1. Clone the Repository
git clone https://github.com/bigxxby/go-test-mobydev
cd go-test-mobydev

### 2. Run Locally

To run the application locally without Docker, execute the following commands:

go run ./cmd/web

This will start the server locally, and you can access the application at `http://localhost:8080` in your web browser.

### 3. Run with Docker

If you prefer to run the application in a Docker container, follow these steps:

Build the Docker image:

docker build -t myApp .

Run the Docker container:

docker run -p 8080:8080 myApp



Now you can access the application at `http://localhost:8080` in your web browser.

### Admin Panel

An admin panel is available with the following credentials:

- Email: admin@example.com
- Password: 12345678Aa#



