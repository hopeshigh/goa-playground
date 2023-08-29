package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("todo", func() {
	Title("Todo Service")
	Description("Everyone learns with todos")
	Server("todo", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("todo", func() {
	Description("Operations for Todos")

	HTTP(func() {
		Path("/")
	})

	Method("create", func() {
		Payload(func() {
			Attribute("title", String, "Todo title")
			Attribute("description", String, "Todo description")
			Required("title", "description")
		})

		Result(String)

		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("complete", func() {
		Payload(func() {
			Attribute("id", String, "Todo id")
			Required("id")
		})

		HTTP(func() {
			POST("/{id}")
			Response(StatusCreated)
		})
	})

	Method("view", func() {
		Payload(func() {
			Attribute("id", String, "Todo id")
			Required("id")
		})

		Result(TodoResult)

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("list", func() {
		Result(CollectionOf(TodoResult))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

var TodoResult = ResultType("TodoResult", func() {
	Attribute("ID", Int, "Unique identifier for the Todo")
	Attribute("Title", String, "Title of the Todo")
	Attribute("Description", String, "Description of the Todo")
	Attribute("Status", String, "Status of the Todo")
	Required("ID", "Title", "Description", "Status")
})
