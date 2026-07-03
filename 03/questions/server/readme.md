This is a great addition. Dockerizing a Go application using multi-stage builds is a critical "Day 1" skill for any backend developer. It keeps your final image tiny by excluding the source code and build tools.

Here is the updated assignment with the Docker requirement included.

***

# Module 3: The Native HTTP Server
## Assignment
Build a performant web server using **only** Go's standard library. You are strictly forbidden from using third-party web frameworks. You must implement custom logging middleware and containerize your application using a multi-stage Dockerfile.

📁 **Project Architecture**
```text
project/           
├── main.go              # Server setup and middleware logic
├── static/
│   └── index.html       # Simple HTML file
├── Dockerfile           # Multi-stage build definition
└── go.mod
```

### Requirements

1.  **Endpoints:**
    *   `GET /json`: Return JSON `{"message": "this is json"}` with the correct `Content-Type`.
    *   `GET /string`: Return text `"i am a string"`.
    *   `GET /html`: Serve `static/index.html` using `http.ServeFile`.

2.  **Logging Middleware:**
    *   Implement a wrapper to log: `[TIME] METHOD PATH STATUS`.
    *   Example: `[2023-10-27 10:00:00] GET /json 200`

3.  **Containerization:**
    *   Write a **Multi-stage `Dockerfile`**.
        *   **Stage 1 (Builder):** Use an official `golang` image to compile your binary.
        *   **Stage 2 (Runner):** Use a minimal image (like `alpine` or `scratch`) to copy the binary and the `static/` folder.
    *   The container must expose port `8080`.

---

### 💡 Implementation Tips

*   **Middleware Pattern:** Create a wrapper function that takes an `http.Handler` and returns an `http.Handler`. To capture the status code, create a simple struct that implements `http.ResponseWriter` and intercepts the `WriteHeader` call.
*   **Dockerfile Multi-Stage Strategy:**
    ```dockerfile
    # Stage 1: Build the binary
    FROM golang:1.23-alpine AS builder
    WORKDIR /app
    COPY . .
    RUN go build -o server main.go

    # Stage 2: Final minimal image
    FROM alpine:latest
    WORKDIR /root/
    COPY --from=builder /app/server .
    COPY --from=builder /app/static ./static
    EXPOSE 8080
    CMD ["./server"]
    ```
*   **Static Files in Docker:** Ensure your code references the static files using a relative path that works *inside* the container’s filesystem, not your local machine's path.

### 📝 Deliverable
Submit your `main.go`, `static/index.html`, and `Dockerfile`. Your mentor will verify that the application runs locally and that the Docker image is built efficiently.