# Note Taking App with Echo

## Assignment
Build a **server-side rendered note-taking web application** using the **Echo framework**.

The application must support:

- signup
- login
- authentication
- creating notes
- deleting notes
- viewing a note by its ID

Use **PostgreSQL** as the backend database.

You only need to provide **Docker for PostgreSQL**.  
There is **no need to dockerize the whole Go project**.

---

# Main Features

## 1. Home Page `/`
When a user visits `/`:

- if the user is **authenticated**, show their notes page as an HTML page
- if the user is **not authenticated**, redirect them to `/login`

On this page, the authenticated user must be able to:

- see the list of their notes
- create a new note
- delete a note

The notes should be listed by **name**.

Each note must be viewable at:

```text
/note/:id
```

---

## 2. Login Page `/login`
This page is for existing users.

The user submits:

- username
- password

Behavior:

- if the account exists and the password is correct, log the user in and redirect to `/`
- if the username does not exist, tell the user to go to `/signup`
- if the password is incorrect, show an error message

---

## 3. Signup Page `/signup`
This page is for creating a new account.

The user submits:

- username
- password

Behavior:

- create the account
- log the user in after successful signup
- redirect the user to `/`

---

# Technical Requirements

## 1. Framework
- The project **must use Echo**
- Handlers must use `echo.Context`

---

## 2. Rendering
- The app must be **SSR** (server-side rendered)
- Use HTML templates to render pages
- Do not build this as an API-only project

Required pages:

- login page
- signup page
- notes list page
- single note page

---

## 3. Database
You must use **PostgreSQL** for storing data.

The database should store at least:

- users
- notes

You should provide:

- SQL schema or migration
- `docker-compose.yml` or Docker command for PostgreSQL only

You do **not** need to dockerize:
- the Echo app
- Go source code
- templates

---

## 4. Authentication
The application must support login state across requests.

Use:
- cookies
- or session-based authentication

Protected routes must require authentication.

If the user is not logged in and tries to access protected pages, redirect them to:

```text
/login
```

---

## 5. Authorization
A user must only be able to access **their own notes**.

That means:
- user A must not be able to view user B’s note
- user A must not be able to delete user B’s note

If a user tries to access another user’s note, return:
- `404 Not Found`
- or `403 Forbidden`

---

## 6. Password Security
Passwords must **not** be stored in plain text.

Use password hashing such as:

```go
golang.org/x/crypto/bcrypt
```

---

# Required Routes

## Public Routes

### `GET /login`
Render the login page.

### `POST /login`
Process login form:
- check if user exists
- verify password
- create authentication session/cookie
- redirect to `/` on success

---

### `GET /signup`
Render signup page.

### `POST /signup`
Process signup form:
- validate input
- ensure username is unique
- hash password
- create user in PostgreSQL
- log the user in
- redirect to `/`

---

## Protected Routes

### `GET /`
If logged in:
- render the user’s notes page

If not logged in:
- redirect to `/login`

The page should include:
- username
- list of note names
- form to create note
- delete button for each note
- link or button to logout

---

### `GET /note/:id`
Show a single note page.

Only the owner of the note may access it.

---

### `POST /notes`
Create a new note for the logged-in user.

A note should at least contain:
- name
- content

---

### `POST /notes/:id/delete`
Delete a note belonging to the logged-in user.

---

### `POST /logout`
Log the user out and redirect to `/login`.

---

# Data Model

## Users Table
Suggested fields:

- `id`
- `username`
- `password_hash`
- `created_at`

---

## Notes Table
Suggested fields:

- `id`
- `user_id`
- `name`
- `content`
- `created_at`

---

# Page Requirements

## Login Page
Must contain:
- username input
- password input
- login button
- link to signup page
- error message area

---

## Signup Page
Must contain:
- username input
- password input
- signup button
- link to login page
- error message area

---

## Home Page
Must contain:
- welcome message
- note creation form
- list of note names
- delete button for each note
- logout button

Each note name should link to:

```text
/note/:id
```

---

## Single Note Page
Must contain:
- note name
- note content
- back link to `/`

---

# Validation Rules

## Signup
- username must not be empty
- password must not be empty
- username must be unique

## Login
- username must exist
- password must match the hashed password

## Note Creation
- note name must not be empty
- note content must not be empty

---

# Middleware
You should use authentication middleware for protected routes.

The middleware should:
- check whether the user is logged in
- redirect unauthenticated users to `/login`

This middleware should protect:

- `/`
- `/note/:id`
- `/notes`
- `/notes/:id/delete`
- `/logout`

---

# Suggested Project Structure

```text
project/
├── main.go
├── handler/
│   ├── auth.go
│   └── note.go
├── middleware/
│   └── auth.go
├── model/
│   ├── user.go
│   └── note.go
├── repo/
│   ├── user_repo.go
│   └── note_repo.go
├── service/
│   ├── auth_service.go
│   └── note_service.go
├── templates/
│   ├── login.html
│   ├── signup.html
│   ├── index.html
│   └── note.html
├── sql/
│   └── schema.sql
├── docker-compose.yml
└── go.mod
```

---

# Docker Requirement
You only need to run **PostgreSQL in Docker**.

For example, students can provide a `docker-compose.yml` like this:

The Go app can run normally on the host machine.

---

# Expected Behavior Example

## New User
1. user opens `/`
2. redirected to `/login`
3. user goes to `/signup`
4. creates account
5. redirected to `/`
6. creates notes
7. clicks a note name to open `/note/:id`

## Existing User
1. user opens `/login`
2. enters username and password
3. redirected to `/`
4. sees only their own notes
5. creates or deletes notes

---

# Constraints
- must use **Echo**
- must use **PostgreSQL**
- must use **SSR**
- must use authentication
- must hash passwords
- must protect note ownership
- must provide Docker only for PostgreSQL


