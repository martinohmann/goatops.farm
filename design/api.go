package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("goatopsfarm", func() {
	Title("goatops.farm")
	Description("Service for obtaining your daily dose of facts about goats and other creatures.")
	Server("goatopsfarm", func() {
		Host("production", func() {
			Description("public host")
			URI("https://goatops.farm")
		})
		Host("development", func() {
			Description("development host")
			URI("http://localhost:8080")
		})
	})

	cors.Origin("*")
})

var _ = Service("creatures", func() {
	Description("The creatures service provides you with farm creatures and facts about them.")

	Error("bad_request")
	Error("not_found")

	Method("list", func() {
		Result(func() {
			Attribute("creatures", ArrayOf(Creature), "List of creatures")
			Required("creatures")
		})

		HTTP(func() {
			GET("/api/v1/creatures")
			Response(StatusOK)
		})
	})

	Method("get", func() {
		Payload(func() {
			Attribute("name", String, "Name of the creature")
			Required("name")
		})

		Result(func() {
			Attribute("creature", Creature, "The creature")
			Required("creature")
		})

		HTTP(func() {
			GET("/api/v1/creatures/{name}")
			Params(func() {
				Param("name", func() {
					Example("goat")
				})
			})
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("random-facts", func() {
		Payload(func() {
			Attribute("name", String, "Name of the creature")
			Attribute("n", Int, "Number of random facts")
			Required("name")
		})

		Result(func() {
			Attribute("facts", ArrayOf(String), "Random facts about the creature")
			Required("facts")
		})

		HTTP(func() {
			GET("/api/v1/creatures/{name}/random-facts")
			Params(func() {
				Param("n", func() {
					Example(3)
				})
				Param("name", func() {
					Example("goat")
				})
			})
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
			Response("not_found", StatusNotFound)
		})
	})
})

var _ = Service("static", func() {
	Description("Static pages and site assets")

	Files("/", "./static/home.html")
	Files("/openapi.json", "./gen/http/openapi3.json")
	Files("/swagger-ui.html", "./static/swagger-ui.html")
})

var Creature = Type("Creature", func() {
	Attribute("name", String, "Name of the creature")
	Attribute("facts", ArrayOf(String), "Facts about the creature")
	Required("name", "facts")
})
