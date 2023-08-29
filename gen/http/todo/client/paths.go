// Code generated by goa v3.12.4, DO NOT EDIT.
//
// HTTP request path constructors for the todo service.
//
// Command:
// $ goa gen github.com/hopeshigh/goa-playground/design

package client

import (
	"fmt"
)

// CreateTodoPath returns the URL path to the todo service create HTTP endpoint.
func CreateTodoPath() string {
	return "/"
}

// CompleteTodoPath returns the URL path to the todo service complete HTTP endpoint.
func CompleteTodoPath(id string) string {
	return fmt.Sprintf("/%v", id)
}

// ViewTodoPath returns the URL path to the todo service view HTTP endpoint.
func ViewTodoPath(id string) string {
	return fmt.Sprintf("/%v", id)
}

// ListTodoPath returns the URL path to the todo service list HTTP endpoint.
func ListTodoPath() string {
	return "/"
}
