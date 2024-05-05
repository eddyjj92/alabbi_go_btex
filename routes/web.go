package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Web() {

	facades.Route().Static("storage", "./storage/app/public")
	facades.Route().Static("dist", "./public/dist")

	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().File("./public/dist/index.html")
	})
}
