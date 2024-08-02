# Strong Password Recommendation API

This project implements a REST API that calculates the minimum number of steps required to make a given password strong. It's built using Go with the Gin framework, and uses PostgreSQL for logging requests and responses.

## Project Structure

strong_password_api/
│
├── main.go           # Main application file
├── main_test.go      # Unit tests
├── Dockerfile        # Dockerfile for the Go application
├── docker-compose.yml# Docker Compose configuration
├── nginx.conf        # Nginx configuration
└── README.md         # This file

## Requirements

- Docker
- Docker Compose

## Setup and Running

1. Clone the repository:
git clone https://github.com/yourusername/strong_password_api.git
cd strong_password_api

2. Build and run the Docker containers:
docker-compose up --build

3. The API will be available at `http://localhost/api/strong_password_steps`

---------------------------------------------------------------------------------------

## API Usage

Send a POST request to `/api/strong_password_steps` with a JSON body:

```json
{
"init_password": "YourPasswordHere"
}


The API will respond with the number of steps required to make the password strong:

{
  "num_of_steps": 3
}

---------------------------------------------------------------------------------------

Password Strength Criteria
A strong password must meet the following criteria:

Length between 6 and 20 characters (inclusive)
Contains at least 1 lowercase letter, 1 uppercase letter, and 1 digit
Does not contain 3 repeating characters in a row

Running Tests
To run the unit tests:
docker-compose exec app go test -v

---------------------------------------------------------------------------------------

Logging
All requests and responses are logged in the PostgreSQL database.
Development
To make changes to the application:

Modify the Go code in main.go
Rebuild and restart the containers:
docker-compose up --build