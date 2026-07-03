package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/DKeshavarz/goraz/module3/taskmanger/service"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(s service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// --- Handler Actions ---

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		h.writeJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	//TODO
	json.NewEncoder(w).Encode(map[string]string{"message": "Task created successfully"})
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request, id int) {
	//TODO
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request, id int) {
	//TODO
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}

// Helper utility for uniform JSON errors
func (h *TaskHandler) writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
