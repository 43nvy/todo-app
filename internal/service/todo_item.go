package service

import (
	"github.com/43nvy/todo-app"
	"github.com/43nvy/todo-app/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (i *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := i.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return i.repo.Create(listId, item)
}

func (i *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return i.repo.GetAll(userId, listId)
}

func (i *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return i.repo.GetById(userId, itemId)
}

func (i *TodoItemService) Delete(userId, itemId int) error {
	return i.repo.Delete(userId, itemId)
}

func (i *TodoItemService) Update(userId, itemId int, input todo.UpdateItemInput) error {
	return i.repo.Update(userId, itemId, input)
}
