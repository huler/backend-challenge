# Problems

## Codebase Review:

This document outlines bugs, and design decisions that were identified and prioritised in the current codebase.

---

### 1. Tightly Coupled Data and Application Layers

- **What is the problem?**:<br>
  The application’s data access logic is directly embedded within the controller logic in files like [getresults/main.go](./api/getresults/main.go) and [postresults/main.go](./api/postresults/main.go). This tight coupling makes the layers of the application code dependent on each other, reducing modularity and flexibility.

- **Why is it a problem?**:<br>
  Tight coupling between the data and application layers violates the principle of separation of concerns. This makes the codebase harder to test, maintain, and extend. For instance:

  - It complicates unit testing since mocking database interactions becomes difficult.
  - Any change in the data access layer might require changes in business logic.
  - Flexibility to switch to a different database or storage engine is severely limited.
  - The readability and reusability of code are diminished.

- **How would you fix this problem?**:<br>
  Introduce a proper abstraction layer for data operations by separating data access logic into a dedicated repository or data layer. This allows:

  - Controllers to remain focused on business logic.
  - Database operations to be tested and mocked independently.
  - Clear interface contracts between layers, making the code more modular and maintainable.

### 2. Repeated DynamoDB Sessions Created For Each Operation on the Data

- **What is the problem?**: <br>
  A new DynamoDB session is being created for every single read, write, or update operation. This results in redundant and inefficient resource usage throughout the application.

- **Why is it a problem?**:<br>
  Creating a new DynamoDB session per operation:

  - Introduces unnecessary overhead in connection handling.
  - Slows down response times, especially under load.
  - Increases the likelihood of throttling and network-related issues.
  - Makes resource management and debugging more difficult.

- **How would you fix this problem?**:
  Use a single shared session per request lifecycle or encapsulate the session within a reusable struct. This was achieved by introducing a `DynamoDBStore` struct during the modularization of the data layer, which holds the session and provides methods for database interactions. This ensures:
  - Reusability and efficiency of a single session instance.
  - Cleaner and more manageable code through encapsulated operations.
