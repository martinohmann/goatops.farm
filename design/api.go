// Copyright 2021 martinohmann
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
