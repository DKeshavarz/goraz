# GoRaz - Module 3: The Go, The Whale, and The Elephant

Welcome to Module 3! Here we will create our first web server, connect a database to it, and finally use our cute whale (Docker) to ship it all together.

## Topics:

1. **JSON (`encoding/json`)**
   * Marshaling/unmarshaling
   * Struct tags
   * Custom JSON parsing

2. **HTTP & Web Servers (`net/http`)**
   * Handlers and HandlerFunc
   * Routing (Standard library vs. third-party routers like Chi/Gin)
   * Middleware implementation
   * Request/Response lifecycle

3. **RESTful APIs**
   * CRUD operations
   * HTTP status codes
   * Input validation
   * Structured error responses

4. **PostgreSQL**
   * SQL basics (DDL & DML)
   * Schema design and relationships
   * CRUD queries and indexing

5. **Connecting Go to PostgreSQL**
   * `database/sql` package
   * PostgreSQL drivers (Recommended: `jackc/pgx`, or legacy: `lib/pq`)
   * Context and timeouts
   * Transactions (`tx`)

6. **Docker & Containerization**
   * Writing a multi-stage `Dockerfile` for Go
   * Docker Compose configuration
   * Running PostgreSQL in a container
   * Containerizing the Go API and linking services

## Questions & Assignments

All questions and assignments are available on Quera. Complete them in order.

🔗 **Solve all questions here:** [Quera Course Link](https://quera.org/course/add_to_course/course/28671/)

*Note: Please ask your mentor for the course password.*

## Resources to Study:

* **Go Web Examples:** [gowebexamples.com](https://gowebexamples.com/) (Great hands-on snippets for HTTP servers)
* **Official Go Tutorial:** [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
* **PostgreSQL & Go:** [pgx documentation](https://github.com/jackc/pgx)
* **Dockerizing Go:** [Docker Official Guide for Go](https://docs.docker.com/guides/golang/)
