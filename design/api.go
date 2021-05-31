package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("goatopsfarm", func() {
	Title("Goat facts Service")
	Description("Service for obtaining your daily dose of goat facts")
	Server("goatopsfarm", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
		Host("goatops.farm", func() {
			Description("public host")
			URI("https://goatops.farm")
		})
	})
})

var _ = Service("goatfacts", func() {
	Description("The goatfacts service provides you with important facts about goats.")

	cors.Origin("*")

	Method("ListFacts", func() {
		Result(ArrayOf(String))

		HTTP(func() {
			GET("/api/facts")
			Response(StatusOK)
		})
	})

	Method("RandomFacts", func() {
		Payload(func() {
			Field(1, "n", Int, "Number of random facts")
		})

		Result(ArrayOf(String))

		Error("BadRequest")

		HTTP(func() {
			GET("/api/facts/random")
			Params(func() {
				Param("n")
			})
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("Index", func() {
		Result(Bytes)

		HTTP(func() {
			GET("/")
			Response(StatusOK, func() {
				ContentType("text/html")
			})
		})
	})

	Files("/api/openapi.json", "./gen/http/openapi3.json")
	Files("/static/{*path}", "./static")
})
