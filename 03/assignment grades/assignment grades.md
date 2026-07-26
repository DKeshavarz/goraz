# JSON

|Part|Points|
|---|--:|
|Correct functionality|8|
|Code structure & readability|3|
|Error handling|3|
|Edge cases|2|
|Tests|2|
|Documentation|2|
|**Total**|**20**|

# DB
| Part                                                           | Points |
| -------------------------------------------------------------- | -----: |
| **Schema design (`schema.sql`)**                               |  **9** |
| • `students` and `books` tables                                |      2 |
| • M:M `loved_books` relationship                               |      2 |
| • 1:1 `borrows` relationship with correct `UNIQUE` constraints |      3 |
| • Self-referencing M:M `friends` relationship                  |      2 |
| **Docker Compose setup (`docker-compose.yml`)**                |  **4** |
| • PostgreSQL container configuration                           |      2 |
| • Automatic schema and data initialization                     |      2 |
| **Test data (`data.sql`)**                                     |  **1** |
| • Valid sample data (including required `golang` student)      |      1 |
| **SQL queries (`query.sql`)**                                  |  **6** |
| • List all students                                            |      1 |
| • Find books loved by `golang`                                 |      1 |
| • Find friends of `golang`                                     |      2 |
| • Find borrowed books with borrower names                      |      2 |
| **Total**                                                      | **20** |

# Web Server
| Part                                                                   | Points |
| ---------------------------------------------------------------------- | -----: |
| **HTTP endpoints (`main.go`)**                                         |  **8** |
| • `GET /json` with correct JSON and `Content-Type`                     |      3 |
| • `GET /string`                                                        |      2 |
| • `GET /html` serving `static/index.html`                              |      3 |
| **Logging middleware**                                                 |  **5** |
| • Middleware implementation (`http.Handler` wrapper)                   |      3 |
| • Correct log format (`[TIME] METHOD PATH STATUS`)                     |      2 |
| **Multi-stage Dockerfile**                                             |  **6** |
| • Correct builder stage (compile Go binary)                            |      2 |
| • Minimal runner stage (`alpine`/`scratch`)                            |      2 |
| • Copies binary & static files, exposes port `8080`, runs successfully |      2 |
| **Project structure & code quality**                                   |  **1** |
| • Clean organization and idiomatic Go                                  |      1 |
| **Total**                                                              | **20** |

# Task manager

|Part|Points|
|---|--:|
|**Project architecture & dependency injection**|**5**|
|• Correct layered architecture (Model → Repository → Service → Delivery)|3|
|• Proper dependency injection from `main.go`|2|
|**Model (`model/task.go`)**|**3**|
|• Correct `Task` struct, JSON tags, and `CreatedAt` type|3|
|**Repository (`repository/task_repo.go`)**|**12**|
|• `FetchAll()`|3|
|• `Create()`|3|
|• `GetByID()`|3|
|• `Delete()`|3|
|**Service (`service/task_service.go`)**|**5**|
|• Correct implementation of service layer|3|
|• Proper interaction with repository & error propagation|2|
|**HTTP Delivery (`delivery/http_handler.go`)**|**15**|
|• `GET /tasks`|3|
|• `POST /tasks`|4|
|• `GET /tasks/{id}`|3|
|• `DELETE /tasks/{id}`|3|
|• All responses are valid JSON with appropriate status codes|2|
|**Database infrastructure**|**7**|
|• `schema.sql` creates the required table correctly|2|
|• `docker-compose.yml` correctly runs PostgreSQL|3|
|• Go application successfully connects to the Dockerized database|2|
|**Code quality & error handling**|**3**|
|• Clean, idiomatic Go, proper package organization, and consistent error handling|3|
|**Total**|**50**|