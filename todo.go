package todoapi

import (
	"context"
	"errors"
	"fmt"
	"log"

	models "github.com/hopeshigh/goa-playground/db/models"
	todo "github.com/hopeshigh/goa-playground/gen/todo"
	"gorm.io/gorm"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct {
	logger *log.Logger
	db     *gorm.DB
}

// NewTodo returns the todo service implementation.
func NewTodo(logger *log.Logger, db *gorm.DB) todo.Service {
	return &todosrvc{logger, db}
}

// Create implements create.
func (s *todosrvc) Create(ctx context.Context, p *todo.CreatePayload) (res string, err error) {
	newTodo := &models.Todo{
		Title:       p.Title,
		Description: p.Description,
		Status:      "Pending",
	}
	if err := s.db.Create(&newTodo).Error; err != nil {
		log.Println("Failed to create a new todo:", err)
		return "", err
	}

	return fmt.Sprintf("Todo Created with ID: %d", newTodo.ID), nil
}

// Complete implements complete.
func (s *todosrvc) Complete(ctx context.Context, p *todo.CompletePayload) (err error) {
	db := s.db.Model(&models.Todo{}).Where("id = ? AND status = ?", p.ID, "Pending").Updates(models.Todo{Status: "Complete"})
	if db.Error != nil {
		log.Printf("Failed to update the todo with ID: %v. Error: %v", p.ID, db.Error)
		return db.Error
	}
	if db.RowsAffected == 0 {
		log.Printf("No todo found with ID: %v in 'Pending' status.", p.ID)
		return errors.New("todo not found or already complete")
	}
	return nil
}

// View implements view.
func (s *todosrvc) View(ctx context.Context, p *todo.ViewPayload) (res *todo.Todoresult, err error) {
	var t models.Todo

	if err := s.db.First(&t, p.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Handle not found error
			log.Println("Todo not found")
			return nil, err
		} else {
			// Handle other errors
			log.Println("Failed to retrieve the todo:", err)
			return nil, err
		}
	}

	res = &todo.Todoresult{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
	}
	s.logger.Print("todo.view")
	return res, nil
}

// List implements list.
func (s *todosrvc) List(ctx context.Context) (res todo.TodoresultCollection, err error) {
	var todos []models.Todo

	if err := s.db.Find(&todos).Error; err != nil {
		log.Println("Failed to retrieve todos:", err)
		return nil, err
	}

	res = make(todo.TodoresultCollection, len(todos))

	for i, t := range todos {
		res[i] = &todo.Todoresult{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
		}
	}

	return res, nil
}
