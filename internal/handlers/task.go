package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/Dharsansivaloganathan/task-manager/internal/database"
    "github.com/Dharsansivaloganathan/task-manager/internal/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    _, err := database.DB.Exec("INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)",
        task.Title, task.Description, "pending")

    if err != nil {
        http.Error(w, "Error creating task", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}
