package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuizHandler(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		question string
		expected string
	}{
		{"1", "The CBI stands for California Bureau of Investigation."},
		{"2", "Patrick Jane has a keen sense of observation and deduction."},
		{"3", "Invalid question number."},
	}

	for _, tt := range tests {
		t.Run(tt.question, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/quiz?question="+tt.question, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Contains(t, w.Body.String(), tt.expected)
		})
	}
}

func TestIndexHandler(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "The Mentalist Quiz")
}
