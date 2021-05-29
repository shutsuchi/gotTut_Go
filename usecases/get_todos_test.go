package usecases_test

import (
	"fmt"
	"testing"
	"github.com/gomagedon/expectate"
	"github.com/shutsuchi/goTut_Go/models"
	"github.com/shutsuchi/goTut_Go/usecases"
)

var dummyTodos = []models.Todo{
	{
		Title:       "todo 1",
		Description: "description of todo 1",
		IsCompleted: true,
	},
	{
		Title:       "todo 2",
		Description: "description of todo 2",
		IsCompleted: true,
	},
	{
		Title:       "todo 3",
		Description: "description of todo 3",
		IsCompleted: true,
	},
}

type BadTodosRepo struct{}

func(BadTodosRepo) GetAllTodos() ([]models.Todo, error) {
	return nil, fmt.Errorf("something went wrong")
}

type MockTodosRepo struct{}

func(MockTodosRepo) GetAllTodos() ([]models.Todo, error) {
	return dummyTodos, nil
}

func TestGetTodos(t *testing.T) {
	t.Run("Returns ErrInternal when TodosRepository returns error", func(t *testing.T) {
		expect := expectate.Expect(t)

		repo := new(BadTodosRepo)

		todos, err := usecases.GetTodos(repo)

		expect(err).ToBe(usecases.ErrInternal)
		// if err != usecases.ErrInternal {
		// 	t.Fatalf("expected ErrInternal; Got: %v", err)
		// }
		if todos != nil {
			t.Fatalf("Expected todos to be nil; Got: %v", todos)
		}
		// expect(todos).ToBe(nil)
	})

	// Test
	t.Run("Returns todos from TodosRepository", func(t *testing.T) {
		expect := expectate.Expect(t)

		repo := new(MockTodosRepo)

		todos, err := usecases.GetTodos(repo)

		expect(err).ToBe(nil)
		expect(todos).ToEqual(dummyTodos)
	})
}