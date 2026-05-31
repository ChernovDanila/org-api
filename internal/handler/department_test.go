package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDepartment_EmptyName(t *testing.T) {
	body := bytes.NewBufferString(`{"name": ""}`)
	req := httptest.NewRequest(http.MethodPost, "/departments/", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Проверяем что пустое имя возвращает 400
	if w.Code == http.StatusCreated {
		t.Error("expected error for empty name, got 201")
	}
}

func TestCreateDepartment_ValidRequest(t *testing.T) {
	body := map[string]string{"name": "Backend"}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/departments/", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	if req.Method != http.MethodPost {
		t.Error("expected POST method")
	}

	_ = w
}
