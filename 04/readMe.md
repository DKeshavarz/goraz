# GoRaz - Module 4: Sessions, Tokens, and Server-Side Pages

Welcome to Module 4! Here we move beyond basic APIs and learn how real web applications handle users, protected routes, and rendered HTML pages. In this module, students will build authentication flows, write middleware, and create server-side rendered pages using Go.

## Topics:

### Echo Framework

- Introduction to Echo and its request lifecycle
- Routing and route groups in Echo
- Handling requests and responses with echo.Context
- Rendering HTML pages in Echo
- Using middleware in Echo
- Implementing authentication and protected routes in Echo
- Comparing net/http and Echo for web development

### Server-Side Rendering (SSR)

- Rendering HTML with `html/template`
- Template layouts and reusable components
- Passing dynamic data from handlers to templates
- Handling forms with `GET` and `POST`
- Displaying validation and error messages in rendered pages

### Authentication & Authorization

- Authentication vs authorization
- User signup and login flow
- Password hashing with `bcrypt`
- Logout flow
- Protecting routes for authenticated users
- Role-based authorization basics (`user`, `admin`)

### JWT (JSON Web Token)

- What JWT is and when to use it
- JWT structure: header, payload, signature
- Creating and validating JWTs in Go
- JWT vs session/cookie-based authentication
- Storing JWT in cookies vs headers
- Common JWT security concerns

### Cookies, Sessions, and State Management

- What cookies are and how they work
- Setting, reading, and deleting cookies in Go
- Session-based authentication
- Storing user state securely
- Flash messages and basic user session flow

### Middleware

- What middleware is in Go web applications
- Writing custom middleware with `net/http`
- Logging middleware
- Authentication middleware
- Authorization middleware
- Middleware chaining

### Project Structure

- Organizing code into `handlers`, `services`, `repositories`, and `templates`
- Keeping business logic out of handlers
- Reusing middleware and helper functions
- Structuring SSR and auth code in a maintainable way

## Questions & Assignments

All questions and assignments are available on Quera. Complete them in order.


## Resources to Study

- Go `html/template` documentation
- Go `net/http` documentation
- Alex Edwards articles on sessions and web development in Go
- JWT introduction and best practices
- OWASP cheatsheets for authentication and session management

