package design

import . "goa.design/goa/v3/dsl"

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

	Method("get-fact", func() {
		Payload(func() {
			Field(1, "id", String, "ID of the fact")
			Required("id")
		})

		Result(Fact)

		Error("NotFound")
		Error("BadRequest")

		HTTP(func() {
			GET("/facts/{id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("list-facts", func() {
		Result(ArrayOf(Fact))

		HTTP(func() {
			GET("/facts")
			Response(StatusOK)
		})
	})

	Method("get-random-fact", func() {
		Result(Fact)

		Error("NotFound")

		HTTP(func() {
			GET("/facts/random")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Files("/swagger.json", "./gen/http/openapi.json")
	Files("/openapi.json", "./gen/http/openapi3.json")
})

var Fact = Type("Fact", func() {
	Description("A fact about goats.")
	TypeName("Fact")

	Attribute("id", String, "A unique ID", func() {
		Example("123abc")
		MinLength(1)
		MaxLength(255)
	})
	Attribute("text", String, "Fact text", func() {
		Example("Goats will not rick-roll you")
	})

	Required("id", "text")
})
