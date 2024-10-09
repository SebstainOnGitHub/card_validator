# Credit Card Number Validator API

## Overview

This is a simple Go-based HTTP server that provides two endpoints to validate credit card numbers. It checks the card number against a checksum and optionally suggests the card's provider.

## Endpoints

### 1. `/check`
- **Method**: `GET`
- **Description**: Validates a credit card number by computing the check digit and comparing it to the real check digit.
- **Request Query**: A card number passed via URL query string.
- **Response**:
  - **200 OK**: The card number is valid.
  - **400 Bad Request**: Invalid card number or an incorrect request method.

### 2. `/check_suggest`
- **Method**: `GET`
- **Description**: Validates a credit card number and suggests the provider if valid. If the card number belongs to an industry-related card (not a standard credit card provider), it returns a different message.
- **Request Query**: A card number passed via URL query string.
- **Response**:
  - **200 OK**: The card number is valid, and the provider (or industry) is returned.
  - **400 Bad Request**: Invalid card number or an incorrect request method.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/sebstainsgit/card-validator-api.git
   cd card-validator-api

2. **Build the project**: Ensure Go is installed on your machine, then run:
   ```bash
   go build -o card-validator

3. **Run the server**:
   ```bash
   ./card-validator
4. The server will run on port 8080!
