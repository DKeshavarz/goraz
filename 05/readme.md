# GoRaz - Module 5: Production Features & Deployment

Welcome to Module 5! In this final module, students will transition from building local applications to developing production-ready features and deploying their projects to a live server.

## Topics

### 1. Search, Filtering, and Pagination
- Implementing pagination using `LIMIT` and `OFFSET` in SQL queries
- Searching database records using SQL pattern matching (`LIKE`)
- Filtering content dynamically based on URL query parameters
- Sorting results by different fields (e.g., date, title)
- Optimizing database queries for search operations

### 2. File Uploads and Media Management
- Handling multipart form uploads (`multipart/form-data`) in Echo
- Validating uploaded files (size limits, MIME types, and extension checks)
- Storing files securely on the server and generating unique filenames
- Preventing directory traversal and security exploits via filenames
- Serving uploaded files safely as static assets using Echo static middleware

### 3. Roles and Permissions
- Designing and implementing Role-Based Access Control (RBAC)
- Differentiating permissions between roles (e.g., Administrator, Author, Reader)
- Restricting route access using custom Echo middleware
- Enforcing resource ownership (preventing users from editing or deleting others' data)
- Returning appropriate HTTP status codes (`403 Forbidden` vs `404 Not Found`)

### 4. Configuration Management
- Implementing the 12-Factor App methodology for application config
- Separating code from configuration using environment variables
- Working with `.env` files for local development
- Structuring configurations using libraries like `cleanenv` or `viper`
- Managing environment-specific settings (Development vs. Production)

### 5. Deployment
- Compiling Go binaries for target Linux environments
- Setting up and securing a Linux VPS (Virtual Private Server)
- Managing application processes in background using `systemd` services
- Configuring Caddy or Nginx as a reverse proxy for port forwarding
- Enabling free SSL/HTTPS certificates with Let's Encrypt
- Managing production logs and updating running applications

---

## Questions & Assignments

All questions and assignments are available on Quera. Complete them in order.
