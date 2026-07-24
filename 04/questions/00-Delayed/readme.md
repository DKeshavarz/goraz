# Delayed Home Page with Echo

## Assignment
Build a small web application using the **Echo framework**.

The application has only **one page**: the home page.  
Whenever a user opens the page, the server must wait for a **random delay** before sending the response.

The delay must be randomly chosen from the interval:

$$
(A, B]
$$

That means:
- greater than `A`
- less than or equal to `B`

Users must also be able to update the delay interval using a **POST request**.

In addition, the application must include a **middleware** that logs the request time in the terminal.

---

## Requirements

### 1. Framework
- The project **must use Echo**
- Request handlers must use `echo.Context`

### 2. Home Page
- The application must have a home page at:

```text
/
```

- Visiting this page with a **GET** request should:
  1. choose a random delay from the current interval `(A, B]`
  2. wait for that amount of time
  3. return the HTML page

- The page should display:
  - the current values of `A` and `B`
  - the delay used for the current response
  - a simple form for updating the interval

---

### 3. Changing the Delay Interval
- The home page must include a form that sends a **POST** request
- The POST request should update the current values of `A` and `B`

For example:

```text
POST /
```

with form fields:
- `a`
- `b`

After updating the interval:
- the new values must be stored in application memory
- future GET requests must use the updated interval

---

### 4. Delay Rules
- The delay must be random in the interval:

$$
(A, B]
$$

- The program must validate the interval:
  - `A` and `B` must be valid numbers
  - `A < B`
  - both should be non-negative

If the input is invalid:
- return a proper error message
- or re-render the page with a validation message

---

### 5. Middleware
Create a middleware that logs the time of each request in the terminal.

At minimum, the log should include:
- request method
- request path
- request time

Example:

```text
[2026-07-24 14:21:03] GET /
[2026-07-24 14:21:08] POST /
```

---

## Suggested Project Structure

```text
project/
├── main.go
├── handler/
│   └── home.go
├── middleware/
│   └── logger.go
├── templates/
│   └── index.html
└── go.mod
```

---

## Route Behavior

### `GET /`
- waits for a random delay in `(A, B]`
- renders the home page

### `POST /`
- reads `a` and `b` from the submitted form
- validates them
- updates the delay interval in memory
- returns a response or redirects back to `/`

---

## Technical Notes

- Use Echo for routing and handlers
- Use `html/template` for rendering the page
- Use `math/rand` for random delay generation
- Use `time.Sleep(...)` to apply the delay
- Store the current interval in memory

---

## Important Implementation Detail

If `A` and `B` are stored globally or shared across requests, students should be careful about concurrent access.

If you want to keep the assignment simple, you can allow a shared in-memory variable.  
If you want to make it slightly stronger, require safe access using:
- `sync.Mutex`
- or another synchronization method

---

## Example Page Content

The home page can show:

- title: `Delayed Home Page`
- current interval
- actual delay used for this response
- form fields for `A` and `B`
- submit button

---

## Validation Examples

### Valid
- `A = 1`, `B = 5`
- `A = 0`, `B = 2`

### Invalid
- `A = 5`, `B = 5`
- `A = 7`, `B = 3`
- negative values
- non-numeric values




