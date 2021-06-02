package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("goatopsfarm", func() {
	Title("goatops.farm")
	Description("Service for obtaining your daily dose of facts about goats and other creatures")
	Server("goatopsfarm", func() {
		Host("goatops.farm", func() {
			Description("public host")
			URI("https://goatops.farm")
		})
		Host("localhost", func() {
			Description("development host")
			URI("http://localhost:8080")
		})
	})

	cors.Origin("*")
})

var _ = Service("facts", func() {
	Description("The facts service provides you with important facts about goats and other creatures.")

	Method("list", func() {
		Result(ArrayOf(String))

		HTTP(func() {
			GET("/api/v1/facts")
			Response(StatusOK)
		})
	})

	Method("list-random", func() {
		Payload(func() {
			Field(1, "n", Int, "Number of random facts")
		})

		Result(ArrayOf(String))

		Error("BadRequest")

		HTTP(func() {
			GET("/api/v1/facts/random")
			Params(func() {
				Param("n", func() {
					Example(10)
				})
			})
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
		})
	})
})

var _ = Service("static", func() {
	Description("Static pages and site assets")

	Files("/", "./static/home.html")
	Files("/openapi.json", "./gen/http/openapi3.json")
	Files("/swagger-ui.html", "./static/swagger-ui.html")
})
