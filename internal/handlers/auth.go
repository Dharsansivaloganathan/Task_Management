package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"

    "github.com/Dharsansivaloganathan/task-manager/internal/database"
    "github.com/Dharsansivaloganathan/task-manager/internal/models"
    "github.com/Dharsansivaloganathan/task-manager/pkg/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    _, err = database.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, hashedPassword)
    if err != nil {
        http.Error(w, "Error registering user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    var storedPassword string

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    err := database.DB.QueryRow("SELECT password FROM users WHERE username=$1", user.Username).Scan(&storedPassword)
    if err == sql.ErrNoRows || !utils.CheckPasswordHash(user.Password, storedPassword) {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateToken(user.Username)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
