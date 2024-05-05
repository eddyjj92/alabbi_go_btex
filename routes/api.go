package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	conversionController := controllers.NewConversionController()
	facades.Route().Group(func(router route.Router) {
		router.Post("/api/conversion/upload", conversionController.Upload)
		router.Post("/api/conversion/start", conversionController.Start)
		router.Post("/api/conversion/stop", conversionController.Stop)
		router.Get("/api/conversion/remove_logs", conversionController.RemoveLogs)
	})

	explorerController := controllers.NewExplorerController()
	facades.Route().Group(func(router route.Router) {
		router.Post("/api/open_folder", explorerController.Index)
		router.Post("/api/output_file", explorerController.OutputFile)
	})

	processController := controllers.NewProcessController()
	facades.Route().Group(func(router route.Router) {
		router.Get("/api/processes", processController.Index)
	})

}
