# Problems

## Codebase Review:

This document outlines bugs, and design decisions that were identified and prioritised in the current codebase.

---

## Assumptions:

- It has been assumed that the **project structure is not to be changed significantly**. All improvements and fixes are applied within the current modular boundaries and file organization.
- It has been assumed that the **external implementations such as AWS DynamoDB configuration, IAM permissions, and environment variable setup (e.g., `TableName`) are correctly provisioned** and functioning as expected.

## Part 1 Solution:

This document outlines **7 major bugs, design flaws, and architectural improvements** observed in the current implementation. Each issue is discussed in detail with an explanation of the problem, its implications, and a proposed or implemented fix.

While the focus has been on addressing significant concerns affecting data integrity, maintainability, and developer experience, it should be noted that **a few minor issues have not been explicitly documented**. These include non-critical edge cases, stylistic inconsistencies, and minor logging enhancements. This decision has been made to keep the document **concise, focused, and actionable**.

## Part 2 Solution:

The submitted pull request provides complete fixes and improvements for the **first five issues** identified in the previous section. Specifically, it addresses:

1. **Tightly Coupled Data and Application Layers**  
   → Refactored to decouple business logic from data access logic using the `DynamoDBStore` abstraction.

2. **Repeated DynamoDB Sessions Created For Each Operation on the Data**  
   → Centralized session creation and reused a single DynamoDB client instance across operations.

3. **Non-Atomic Database Operations**  
   → Replaced multiple individual writes with a single `TransactWriteItems` call for atomic behavior.

4. **Missing Error Handling in Concurrent Data Fetch**
   → Introduced an error channel to capture and log errors from goroutines during concurrent department-wise data retrieval.

5. **Potential nil Dereference in `GetTotalResponseCount`**  
   → Added proper nil checks before unmarshalling the response to prevent runtime panics.


## Problems Identified
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

### 3. Non-Atomic Database Operations

- **What is the problem?**:  
  Database operations that should be logically atomic — such as recording a user's survey response and incrementing a global response counter — were implemented as two separate calls: `PutItem` for the survey data and `UpdateItem` for the count. This leads to potential race conditions and data inconsistencies if one operation succeeds while the other fails.

- **Why is it a problem?**:  
  Executing these operations separately introduces the risk of partial writes:

  - If the response is stored but the count is not updated, analytics and metrics become unreliable.
  - If the count is updated but the response fails, it leads to inflated counts.
  - These inconsistencies are hard to detect and correct after the fact.

- **How would you fix this problem?**:  
  Use DynamoDB’s `TransactWriteItems` to group both the `PutItem` and `UpdateItem` into a single atomic transaction. This ensures that either both operations succeed together, or neither is applied.

### 4. Missing Error Handling in Concurrent Data Fetch

- **What is the problem?**:  
  When using goroutines to concurrently fetch survey responses for various departments in [getresults/main.go](./api/getresults/main.go), errors from individual goroutines were not properly captured or handled. This could cause failures to go unnoticed and lead to incomplete or misleading results being returned.

- **Why is it a problem?**:  
  Without capturing and reporting errors from goroutines:

  - Failures in data retrieval (e.g., failed scans, unmarshalling) are silently ignored.
  - Results can appear as empty or incorrect without any indication of failure.
  - Debugging becomes significantly harder since the main thread has no visibility into errors within the goroutines.

- **How would you fix this problem?**:  
  Introduce an error channel (`chan error`) and pass it to each goroutine. Each goroutine sends either `nil` or an error back to the channel. After all goroutines complete, errors are collected and logged centrally. This provides reliable visibility into all parallel execution paths and ensures system robustness.

### 5. Potential Nil Dereference in `GetTotalResponseCount`

- **What is the problem?**:  
  The function `GetTotalResponseCount` retrieves a specific item from DynamoDB and then attempts to unmarshal the result into a Go struct. However, if the requested item does not exist the subsequent call to unmarshall can cause a panic due to a nil dereference.

- **Why is it a problem?**:  
  If `data.Item` is `nil`, it indicates that the requested item does not exist in DynamoDB. Attempting to unmarshal a `nil` value results in a runtime panic, which:

  - Crashes the program or Lambda execution.
  - Makes the service unreliable.
  - Can lead to increased operational cost and degraded user experience if not handled properly.

- **How would you fix this problem?**:  
  Add a check immediately after retrieving the item to verify that `data.Item` is not `nil`. If it is `nil`, return a descriptive error before attempting any unmarshalling.

### 6. Inconsistent Error Handling and Logging in Controller

- **What is the problem?**:  
  The controllers use inconsistent error handling practices:

  - It prints errors using `fmt.Printf` but returns generic HTTP 500 responses without context.
  - There's no structured or centralized logging, which limits observability and makes debugging production issues harder.

- **Why is it a problem?**:  
  Inconsistent and informal error handling has multiple consequences:

  - Logs lack traceability: developers can't correlate logs with specific failures or requests.
  - Monitoring systems cannot detect specific error types.

- **How would you fix this problem?**:

  - Use structured logging instead of `fmt.Printf`, possibly integrating with a logging library that supports severity levels and context propagation.
  - Add more granular HTTP response codes (e.g., 400 for client validation errors, 500 for internal errors).
  - Include request identifiers or metadata in logs for traceability.

  ### 7. Missing Request Payload Validation (Email and Department)

- **What is the problem?**:  
  The incoming request payload is not strictly validated for expected formats or values. The check only ensures that `Email` and `Department` fields are non-empty and that `Results` contains exactly four integers. However, it does not:

  - Validate whether `Email` has a proper format.
  - Verify if `Department` matches a predefined list of acceptable departments.
  - Ensure individual integers in `Results` fall within a valid range (e.g., 1–5).

- **Why is it a problem?**:  
  Accepting malformed or invalid data can result in:

  - Corrupted or unusable records in the database.
  - Skewed analytics and reporting.
  - Increased risk of unexpected behavior or runtime errors.
  - Potential security vulnerabilities (e.g., injection, abuse via malformed inputs).
    This also forces downstream systems (analytics, dashboards) to account for poor-quality input.

- **How would you fix this problem?**:  
  Implement stricter request validation logic:
  - Use regex to verify `Email` format.
  - Check `Department` against a controlled list (e.g., `"HR"`, `"QA"`, etc.).
  - Validate `Results` to ensure all values are within an expected numeric range.
  - Return HTTP 400 (Bad Request) if the payload is malformed.
    Optionally, encapsulate validation logic in a reusable function for better testability and modularity.
