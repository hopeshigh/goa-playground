// Code generated by goa v3.12.4, DO NOT EDIT.
//
// todo service
//
// Command:
// $ goa gen github.com/hopeshigh/goa-playground/design

package todo

import (
	"context"

	todoviews "github.com/hopeshigh/goa-playground/gen/todo/views"
)

// Operations for Todos
type Service interface {
	// Create implements create.
	Create(context.Context, *CreatePayload) (res string, err error)
	// Complete implements complete.
	Complete(context.Context, *CompletePayload) (err error)
	// View implements view.
	View(context.Context, *ViewPayload) (res *Todoresult, err error)
	// List implements list.
	List(context.Context) (res TodoresultCollection, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "todo"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"create", "complete", "view", "list"}

// CompletePayload is the payload type of the todo service complete method.
type CompletePayload struct {
	// Todo id
	ID string
}

// CreatePayload is the payload type of the todo service create method.
type CreatePayload struct {
	// Todo title
	Title string
	// Todo description
	Description string
}

// Todoresult is the result type of the todo service view method.
type Todoresult struct {
	// Unique identifier for the Todo
	ID int
	// Title of the Todo
	Title string
	// Description of the Todo
	Description string
	// Status of the Todo
	Status string
}

// TodoresultCollection is the result type of the todo service list method.
type TodoresultCollection []*Todoresult

// ViewPayload is the payload type of the todo service view method.
type ViewPayload struct {
	// Todo id
	ID string
}

// NewTodoresult initializes result type Todoresult from viewed result type
// Todoresult.
func NewTodoresult(vres *todoviews.Todoresult) *Todoresult {
	return newTodoresult(vres.Projected)
}

// NewViewedTodoresult initializes viewed result type Todoresult from result
// type Todoresult using the given view.
func NewViewedTodoresult(res *Todoresult, view string) *todoviews.Todoresult {
	p := newTodoresultView(res)
	return &todoviews.Todoresult{Projected: p, View: "default"}
}

// NewTodoresultCollection initializes result type TodoresultCollection from
// viewed result type TodoresultCollection.
func NewTodoresultCollection(vres todoviews.TodoresultCollection) TodoresultCollection {
	return newTodoresultCollection(vres.Projected)
}

// NewViewedTodoresultCollection initializes viewed result type
// TodoresultCollection from result type TodoresultCollection using the given
// view.
func NewViewedTodoresultCollection(res TodoresultCollection, view string) todoviews.TodoresultCollection {
	p := newTodoresultCollectionView(res)
	return todoviews.TodoresultCollection{Projected: p, View: "default"}
}

// newTodoresult converts projected type Todoresult to service type Todoresult.
func newTodoresult(vres *todoviews.TodoresultView) *Todoresult {
	res := &Todoresult{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Title != nil {
		res.Title = *vres.Title
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	return res
}

// newTodoresultView projects result type Todoresult to projected type
// TodoresultView using the "default" view.
func newTodoresultView(res *Todoresult) *todoviews.TodoresultView {
	vres := &todoviews.TodoresultView{
		ID:          &res.ID,
		Title:       &res.Title,
		Description: &res.Description,
		Status:      &res.Status,
	}
	return vres
}

// newTodoresultCollection converts projected type TodoresultCollection to
// service type TodoresultCollection.
func newTodoresultCollection(vres todoviews.TodoresultCollectionView) TodoresultCollection {
	res := make(TodoresultCollection, len(vres))
	for i, n := range vres {
		res[i] = newTodoresult(n)
	}
	return res
}

// newTodoresultCollectionView projects result type TodoresultCollection to
// projected type TodoresultCollectionView using the "default" view.
func newTodoresultCollectionView(res TodoresultCollection) todoviews.TodoresultCollectionView {
	vres := make(todoviews.TodoresultCollectionView, len(res))
	for i, n := range res {
		vres[i] = newTodoresultView(n)
	}
	return vres
}
