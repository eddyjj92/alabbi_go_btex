package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type ProcessController struct {
	//Dependent services
}

func NewProcessController() *ProcessController {
	return &ProcessController{
		//Inject services
	}
}

func (r *ProcessController) Index(ctx http.Context) http.Response {

	var processes []models.Process
	facades.Orm().Query().Get(&processes)
	return ctx.Response().Success().Json(http.Json{
		"processes": processes,
	})

}
