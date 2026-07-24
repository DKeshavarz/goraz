
# Poem Learning Platform

## Assignment
Build a small **server-side rendered web application using the Echo framework** that displays poems from a JSON data source.

The program must load poem data at startup and expose routes for browsing and reading poems.

 **Project Architecture**
```text
project/
├── main.go               # Echo server setup and route registration
├── poem/
│   ├── model/
│   │   └── poem.go       # Poem struct definition
│   ├── repo/
│   │   └── json_repo.go  # JSON loading logic
│   ├── service/
│   │   └── poem.go       # Poem business logic
│   └── handler/
│       └── http.go       # Echo handlers
├── templates/
│   ├── index.html        # Poem list page
│   └── poem.html         # Single poem page
├── data/
│   └── poems.json        # Already given, improve it
├── static/
│   └── poems.jpeg        # Already given, improve it
└── go.mod                # Module definition
```

## Requirements

### 1. Echo Framework is Mandatory
- The application **must be implemented using the Echo framework**
- Route registration must be done with Echo
- Request handling must use `echo.Context`
- HTML rendering must be integrated with Echo
- Static files must be served using Echo

### 2. Poem Data Model
Each poem in the JSON file contains:
- `name`
- `lines` (array of strings)
- `picture` (optional)

### 3. Data Loading
- The program must read `data/poems.json` when the application starts
- The poem data must be stored in memory after loading
- The JSON structure can contain poems in any language

### 4. Routes
Your program must provide the following routes:

#### `/`
- Show a list of all poem names
- Each poem name must be clickable
- Clicking a poem name should open the poem page

#### `/poems/:name`
- Show the selected poem page
- Display:
  - poem name
  - poem picture if it exists
  - all poem lines
- Each line must be hidden by default
- When the user clicks a line, that line should be revealed

### 5. Rendering
- Use **SSR** to render both pages
- Use Go templates together with Echo’s renderer
- The homepage and poem page must be rendered dynamically from the loaded poem data

### 6. Poem Lookup
- If a poem with the requested name does not exist, return a proper **404 page** or HTTP 404 response

### 7. Static Files
- Serve images and other static assets using Echo static file support

## Layer Responsibilities

- **model**
  - defines the `Poem` struct

- **repo**
  - reads poem data from the JSON file
  - loads and stores poem data in memory

- **service**
  - provides poem-related logic
  - returns all poems
  - finds a poem by name

- **handler**
  - handles HTTP requests using Echo
  - renders templates
  - returns proper HTTP responses

---

## Technical Constraints

- You **must use Echo**
- You **must keep the layered architecture**
- You **must use SSR**, not a frontend framework
- You may use a small amount of JavaScript only for revealing poem lines
- Do not place business logic directly inside route registration
- The `picture` field must be handled as optional

---

## Tips for implementation
- Use `encoding/json` to load poem data from file
- Use `html/template` with a custom Echo renderer
- Use `c.Param("name")` to read the poem name from the URL
- Consider using URL-safe names or slugs for poems with spaces or special characters
- Use `e.Static("/static", "static")` for serving images
- Return `c.String(404, "poem not found")` or render a custom 404 template when needed




