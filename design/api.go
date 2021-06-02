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
		Result(func() {
			Attribute("facts", ArrayOf(String), "List of facts")
			Required("facts")
		})

		HTTP(func() {
			GET("/api/v1/facts")
			Response(StatusOK)
		})
	})

	Method("list-random", func() {
		Payload(func() {
			Attribute("n", Int, "Number of random facts")
		})

		Result(func() {
			Attribute("facts", ArrayOf(String), "List of random facts")
			Required("facts")
		})

		Error("bad_request", ErrorResult, "Bad request payload")

		HTTP(func() {
			GET("/api/v1/facts/random")
			Params(func() {
				Param("n", func() {
					Example(10)
				})
			})
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
		})
	})
})

var _ = Service("static", func() {
	Description("Static pages and site assets")

	Files("/", "./static/home.html")
	Files("/openapi.json", "./gen/http/openapi3.json")
	Files("/swagger-ui.html", "./static/swagger-ui.html")
})
