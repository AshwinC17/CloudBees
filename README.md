# Train Ticket Reservation System

Welcome to the Train Ticket Reservation System! This Golang and gRPC-based application allows users to purchase and manage train tickets from London to France.

## Overview

This application leverages gRPC APIs for seamless communication, and data is stored in-memory for simplicity.

## Getting Started

Follow these steps to run the application:

1. Clone the repository:
   ```bash
   git clone https://github.com/AshwinC17/CloudBees.git
   cd CloudBees
   ```

2. Run the server:
   ```bash
   go run server/server.go
   ```

3. Run the client:
   ```bash
   go run client/client.go
   ```

## Usage

### 1. Submit a Ticket Purchase

```bash
go run client/client.go submit-ticket --from London --to France --user "John Doe" --email "john.doe@example.com"
```

### 2. View Receipt Details

```bash
go run client/client.go view-receipt --user "John Doe"
```

### 3. View Users and Allocated Seats by Section

```bash
go run client/client.go view-users --section A
```

### 4. Remove a User from the Train

```bash
go run client/client.go remove-user --user "John Doe"
```

### 5. Modify a User's Seat

```bash
go run client/client.go modify-seat --user "John Doe" --new-seat "A12"
```

## Notes

- The application currently supports two train sections: A and B.
- The ticket price is fixed at $20.

Feel free to explore and modify the code to meet your requirements. We look forward to reviewing your results!
