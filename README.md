# PayTaps-task

PayTaps-task is a Golang solution for mocking money transfer transactions between accounts. It provides a simple way to manage account transactions, and it's written in Go.

## Project Structure

The project is organized as follows:

1. `main.go`: Entry point of the application.
2. `database`: Package responsible for managing accounts and transactions.
3. `services`: Business logic for handling transactions.
4. `controllers`: HTTP logic for serving endpoints.

## Endpoints

There are two main endpoints in the project:

1. **List All Accounts**:
   - Endpoint: `GET /api/accounts/listing`

2. **Make a Transaction**:
   - Endpoint: `POST /api/accounts/transaction`
   - Request Body:
     ```json
     {
         "fromId": "3d253e29-8785-464f-8fa0-9e4b57699db9",
         "toId": "17f904c1-806f-4252-9103-74e7a5d3e340",
         "amount": 10.909
     }
     ```
## How to Run App
  ```
  go run .
  ```
## How to Run Tests

To run the tests, you can use the following command:

```
go test ./test

```

## Database Struct
The `database` package includes essential functions:

1. `getAccountById`: Returns a copy of the account struct based on the account ID.
2. `updateBalance`: Updates the balance for an account and returns a copy of it after the update.
3. `GetAccountLock`: Returns a mutex associated with an account to prevent multiple operations simultaneously, which helps resolve race conditions.

## Service Function
The most critical function in this project is service.MakeTransaction. It accepts transaction data, ensures validity, and updates the accounts with new balances based on the provided amounts.

## Future Enhancements

While this application is ready for the task, there are several areas where it can be further enhanced:

1. **Custom Errors:** Consider creating custom error types to provide more context and information about specific errors. Custom error types can improve the clarity of error messages and help with debugging.

2. **Improving Response Shape:** Define a standard response format to make it easier for clients to understand and work with the API. A well-structured response format can enhance the user experience and simplify client integration.

3. **Increasing Tests:** Expand test coverage to cover various scenarios, including edge cases and error conditions. Comprehensive test coverage ensures the reliability and robustness of the application.

4. **Adding a Configuration Module:** Implement a configuration module to centralize application settings for different environments. This can simplify the management of configuration parameters for development, testing, and production environments.

5. **Adding Utility Modules:** Create utility functions for common tasks like data validation and logging. Utility modules can provide reusable functionality that simplifies the development process and promotes code reusability.

## At the end

 **Hope to hear from you soon**
