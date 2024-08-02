# Strong Password API

This project implements a backend service that recommends steps to create a strong password.

## Project Overview

The service provides an API that takes a password as input and returns the minimum number of actions required to make the password strong based on specific criteria.

### API Specification

- **Base URL**: `/api/strong_password_steps`
- **Method**: POST
- **Input**: JSON object with `init_password` field
- **Output**: JSON object with `num_of_steps` field

#### Example Request:
```json
{
  "init_password": "aA1"
}
```

#### Example Response:
```json
{
  "num_of_steps": 3
}
```


This project implements a backend service that recommends steps to create a strong password.

## Project Overview

The service provides an API that takes a password as input and returns the minimum number of actions required to make the password strong based on specific criteria.

### API Specification

- **Base URL**: `/api/strong_password_steps`
- **Method**: POST
- **Input**: JSON object with `init_password` field
- **Output**: JSON object with `num_of_steps` field

#### Example Request:
```json
{
  "init_password": "aA1"
}
```
### Strong Password Criteria

Password length >=6 and < 20 characters.
Contains at least 1 lowercase letter, at least 1 uppercase letter, and at least 1 digit.
Does not contain 3 repeating characters in a row (e.g., "11123" is invalid).

### Tech Stack

Go with Gin framework
Nginx
PostgreSQL
Docker

### Local Deployment
To deploy the project locally:

1.Ensure you have Docker and Docker Compose installed on your system.

2.Clone the repository
```
git clone https://github.com/your-username/your_name_agnos_backend.git
cd your_name_agnos_backend
```

3.Build and start the services using Docker Compose
```
docker-compose up --build
```

4.The API should now be accessible at
`http://localhost:8080/api/strong_password_steps`


### Running Unit Tests
To run the unit tests:

Ensure you're in the project root directory.
Run the following command:
```json
go test ./...
```
This will run all unit tests in the project.

### Additional Features

Request and response logs are stored in the PostgreSQL database.
The code structure is designed for easy maintenance.


### Notes

Passwords should contain only letters, digits, '.' (dot), or '!' (exclamation mark).
Password length should be between 1 and 40 characters inclusive.

For any issues or questions, please open an issue in the GitHub repository.

`This README provides an overview of the project, explains how to deploy it locally, how to run unit tests, and includes key information about the API and its requirements. You may want to adjust some details based on your specific implementation, such as the exact steps for running tests or any additional setup required for the database.`