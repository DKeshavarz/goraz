# Library fucking DB
## Assignment
Design a relational database for a library where students can interact with books and each other. Your schema must enforce the relationships correctly using SQL constraints.

📁 **Project Architecture**
```text
library-db/           
├── docker-compose.yml  # Setup for PostgreSQL service
├── schema.sql          # Table definitions (DDL)
├── data.sql            # Test data population (DML)
└── query.sql           # Your analysis queries
```

### Requirements

1.  **Schema Design (`schema.sql`):**
    *   **Students:** `id, name, email, phone`
    *   **Books:** `id, title, year, pageCount`
    *   **Relationship 1 (M:M):** Create a table to link students who "love" books.
    *   **Relationship 2 (1:1):** Create a table to track borrowed books. Ensure that a student can only borrow one book at a time, and a book can be borrowed by only one student at a time (Use `UNIQUE` constraints).
    *   **Relationship 3 (Self-Referencing M:M):** Create a table to link `students` to `students` as friends.

2.  **Docker Setup (`docker-compose.yml`):**
    *   Set up a `postgres` container.
    *   Ensure the schema and data are automatically applied when the container starts (using `/docker-entrypoint-initdb.d/`).
  
3.  **Test Data (`data.sql`):**
    *   Populate each table with at least 2 records.
    *   Ensure you have a student named `'golang'` in your data to test the queries below.

4.  **Query Implementation (`query.sql`):**
    Write the SQL queries to answer the following requirements:
    *   **List all students.**
    *   **Find books:** Retrieve the titles of all books that the student named `'golang'` has marked as "loved".
    *   **Find friends:** Retrieve the names of all friends of the student named `'golang'`.
    *   **Find loans:** Retrieve all books currently borrowed (return book title and borrower name).

---

### 💡 Implementation Tips

*   **1:1 Relationship:** For the borrowing system, do not just put a column in the `books` table. Create a `borrows` table with columns `student_id` and `book_id`. Apply a `UNIQUE` constraint to `book_id` (so a book can only be borrowed once) and a `UNIQUE` constraint to `student_id` (so a student can only borrow one book at a time).
*   **Self-Referencing:** When joining the `friends` table, you will need to join the `students` table twice (once for the "student" and once for the "friend").
*   **Initialization:** In your `docker-compose.yml`, mount your SQL files to the container's init directory:
    ```yaml
    volumes:
      - ./schema.sql:/docker-entrypoint-initdb.d/01_schema.sql
      - ./data.sql:/docker-entrypoint-initdb.d/02_data.sql
    ```

### 📝 Deliverable
Submit your `docker-compose.yml`, `schema.sql`, `data.sql`, and `query.sql` files. Your mentor will run `docker-compose up` to verify that your schema is valid and your queries return the correct data.