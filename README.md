# Train Ticket Service

This project implements a RESTful API for managing train tickets, allowing users to purchase tickets, view receipts, manage seats, and more.

## Technologies Used

- Java
- Spring Boot
- Maven

## Getting Started

### Prerequisites

- Java 8 or higher
- Maven

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ashwinc17/trainicketservice.git
   ```

2. Navigate to the project directory:

   ```bash
   cd trainticketservice
   ```

3. Build the project using Maven:

   ```bash
   mvn clean install
   ```

4. Run the application:

   ```bash
   java -jar target/trainticketservice.jar
   ```

5. The application should now be running on `http://localhost:8080`.

## API Endpoints

- `POST /tickets/purchase`: Purchase a train ticket.
- `GET /tickets/receipt/{email}`: Get the receipt for a user.
- `GET /tickets/users/{section}`: Get users and their seats in a specific section.
- `DELETE /tickets/remove/{email}`: Remove a user from the train.
- `PUT /tickets/modify-seat`: Modify a user's seat.


## Usage

### Purchase a Ticket

To purchase a ticket, send a POST request to `/tickets/purchase` with the user details and ticket information in the request body.

Example:

```bash
curl -X POST -H "Content-Type: application/json" -d "{\"firstName\": \"John\", \"lastName\": \"Doe\", \"email\": \"john.doe@example.com\"}" "http://localhost:8080/tickets/purchase?from=London&to=France&price=20.0"
```

### View Receipt

To view a receipt, send a GET request to `/tickets/receipt/{email}` with the user's email address as a path variable.

Example:

```bash
curl -X GET http://localhost:8080/tickets/receipt/john.doe@example.com
```

### Modify Seat

To modify a user's seat, send a PUT request to `/tickets/modify-seat` with the user's email address and the new seat number as query parameters.

Example:

```bash
curl -X PUT "http://localhost:8080/tickets/modify-seat?email=john.doe@example.com&newSeat=B5"
```

### Remove User

To remove a user from the train, send a DELETE request to `/tickets/remove/{email}` with the user's email address as a path variable.

Example:

```bash
curl -X DELETE http://localhost:8080/tickets/remove/john.doe@example.com
```
