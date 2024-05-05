package controllers

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"os/exec"
)

type ExplorerController struct {
	//Dependent services
}

func NewExplorerController() *ExplorerController {
	return &ExplorerController{
		//Inject services
	}
}

func (r *ExplorerController) Index(ctx http.Context) http.Response {

	data := ctx.Request().All()

	cmd := exec.Command(
		"explorer",
		[]string{
			fmt.Sprintf("%s", data["ruta"]),
		}...)
	cmd.Run()

	return ctx.Response().Success().Json(http.Json{
		"message": "Abriendo Carpeta",
	})
}

func (r *ExplorerController) OutputFile(ctx http.Context) http.Response {

	data := ctx.Request().All()

	cmd := exec.Command(
		"notepad.exe",
		[]string{
			fmt.Sprintf("%s", data["file"]),
		}...)
	cmd.Run()

	return ctx.Response().Success().Json(http.Json{
		"message": "Abriendo Archivo",
	})
}
