# Clime coding challenge api

This is a simple Go-based API that performs basic math operations such as addition, subtraction, multiplication, and division.

## Requirements

- Go (1.20 recommended)
- Gin Web Framework
- Go-Cache

## How to Compile and Run

Before running the application, make sure you have installed Go in your system.

1. Install the required packages by running:
    ```shell
    go get github.com/gin-gonic/gin
    go get github.com/patrickmn/go-cache
    ```

2. Compile the Go application by running:
    ```shell
    go build main.go
    ```

3. This will generate an executable file in your directory. Run the executable:
    ```shell
    ./main
    ```

The application is now running and will listen for requests at `localhost:80`

## API Usage

The API currently supports four operations:

- Addition
- Subtraction
- Multiplication
- Division

Each operation corresponds to a route:

- `/add?x=<num1>&y=<num2>`
- `/subtract?x=<num1>&y=<num2>`
- `/multiply?x=<num1>&y=<num2>`
- `/divide?x=<num1>&y=<num2>`

Replace `<num1>` and `<num2>` with the numbers you want to operate on.

Example request: `localhost:80/add?x=10&y=20`

The API will respond with a JSON object like this:

```json
{
  "action": "add",
  "x": 10,
  "y": 20,
  "answer": 30,
  "cached": false
}
