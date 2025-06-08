# Problems

## Codebase Review:

This document outlines bugs, and design decisions that were identified and prioritised in the current codebase.

---

### 1. Tightly Coupled Data and Application Layers

- **Files Affected**:
  - [getresults/main.go](./api/getresults/main.go)
  - [postresults/main.go](./api/postresults/main.go)
- **Type**:<br>Design Decision
- **Reason for Selection**:<br>Data access logic is mixed with business logic, violating separation of concerns and making the system harder to test and refactor.
- **Affects**:
  - Developer Experience
  - Long-term Maintenance
- **Solution**:
- **Before/After**:
  1. Session Creation:<br>
     **Before**: Created a new session everytime records are read/added/updated to database.<br>
     **After**: A single session is created to perform end-to-end operations for a controller.
  1. Data Layer Abstraction:<br>
     **Before**: Data layer was tightly coupled with the controller implementation, causing difficulties for unit testing and flexibility for underlying data store.<br>
     **After**: Data layer is decoupled from the controller implementation, allowing unit testing by mocking databases and making it flexible to use different underlying data structures.
