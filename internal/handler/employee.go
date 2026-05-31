package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ChernovDanila/org-api/internal/service"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// Получаем {id} из URL
	parts := strings.Split(r.URL.Path, "/")
	deptID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "invalid department id", http.StatusBadRequest)
		return
	}

	var req struct {
		FullName string  `json:"full_name"`
		Position string  `json:"position"`
		HiredAt  *string `json:"hired_at"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	var hiredAt *time.Time
	if req.HiredAt != nil {
		t, err := time.Parse("2006-01-02", *req.HiredAt)
		if err != nil {
			http.Error(w, "invalid hired_at format, use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		hiredAt = &t
	}

	emp, err := h.service.Create(deptID, req.FullName, req.Position, hiredAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}
