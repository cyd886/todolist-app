package domain

import (
	"testing"
	"time"
)

func TestTodo_Creation(t *testing.T) {
	now := time.Now()
	todo := &Todo{
		ID:          1,
		Title:       "Test Todo",
		Description: "Test Description",
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if todo.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", todo.ID)
	}

	if todo.Title != "Test Todo" {
		t.Errorf("Expected Title to be 'Test Todo', got %s", todo.Title)
	}

	if todo.Description != "Test Description" {
		t.Errorf("Expected Description to be 'Test Description', got %s", todo.Description)
	}

	if todo.Completed != false {
		t.Errorf("Expected Completed to be false, got %t", todo.Completed)
	}

	if !todo.CreatedAt.Equal(now) {
		t.Errorf("Expected CreatedAt to be %v, got %v", now, todo.CreatedAt)
	}

	if !todo.UpdatedAt.Equal(now) {
		t.Errorf("Expected UpdatedAt to be %v, got %v", now, todo.UpdatedAt)
	}
}
