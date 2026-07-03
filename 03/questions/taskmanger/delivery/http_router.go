package delivery

import (
	"net/http"
	"strconv"
	"strings"
)

// 1. Handles /tasks (GET = list, POST = create)
func (h *TaskHandler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		h.ListTasks(w, r)
	case http.MethodPost:
		h.CreateTask(w, r)
	default:
		h.writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// 2. Handles /tasks/{id} (GET = details, DELETE = remove)
func (h *TaskHandler) TaskMemberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from the path (e.g. "/tasks/12" -> "12")
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if idStr == "" {
		h.writeJSONError(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.writeJSONError(w, "Invalid task ID format", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetTask(w, r, id)
	case http.MethodDelete:
		h.DeleteTask(w, r, id)
	default:
		h.writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
