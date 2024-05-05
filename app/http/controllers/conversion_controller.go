package controllers

import (
	"bytes"
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"
	"github.com/goravel/framework/validation"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type ConversionController struct {
	Helper *HelperController
}

func NewConversionController() *ConversionController {
	return &ConversionController{
		NewHelperController(),
	}
}

func (r *ConversionController) Upload(ctx http.Context) http.Response {
	data := ctx.Request().All()
	rules := map[string]string{
		"file": "required",
	}
	messages := validation.Messages(map[string]string{
		"required": "El campo :attribute es requerido.",
	})
	attributes := validation.Attributes(map[string]string{
		"file": "archivo",
	})
	validator, err := facades.Validation().Make(data, rules, messages, attributes)
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"error":     err.Error(),
			"message":   nil,
			"validator": nil,
		})
	} else if validator.Fails() {
		return ctx.Response().Status(400).Json(http.Json{
			"error":     nil,
			"message":   nil,
			"validator": validator.Errors().All(),
		})
	}
	file, err := ctx.Request().File("file")
	filename := fmt.Sprintf("%s", file.GetClientOriginalName())

	extensions := []string{"mp3", "wma", "wav", "ogg", "webm", "mp4", "mpg", "wmv", "avi"}

	if !r.Helper.FindStringIntoSlice(extensions, file.GetClientOriginalExtension()) {
		return ctx.Response().Status(400).Json(http.Json{
			"validator": "Solo se permiten archivos de audio y video.",
		})
	}

	currentTime := time.Now()
	/*tz, _ := time.LoadLocation("America/Havana")*/
	formated := fmt.Sprintf("%s_%s-%s-%s", currentTime.Format(time.DateOnly), strconv.Itoa(currentTime.Hour()), strconv.Itoa(currentTime.Minute()), strconv.Itoa(currentTime.Second()))

	_, err = file.StoreAs("./public/files/"+formated, fmt.Sprintf("%s", filename))
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"error":     err.Error(),
			"message":   nil,
			"validator": nil,
		})
	}

	ruta, _ := os.Getwd()

	return ctx.Response().Success().Json(http.Json{
		"error":     nil,
		"message":   "Archivo importado listo para procesar.",
		"validator": nil,
		"ruta":      fmt.Sprintf("%s", ruta+"\\"+path.Storage()+"\\app\\public\\files\\"+formated+"\\"+filename),
		"filename":  filename,
		"outputDir": fmt.Sprintf("%s", ruta+"\\"+path.Storage()+"\\app\\public\\files\\"+formated),
		"extension": file.GetClientOriginalExtension(),
		"folder":    formated,
	})
}

func (r *ConversionController) Start(ctx http.Context) http.Response {
	data := ctx.Request().All()
	rules := map[string]string{
		"input":         "required",
		"output_format": "required",
		"output_dir":    "required",
	}
	messages := validation.Messages(map[string]string{
		"required": "El campo :attribute es requerido.",
	})
	attributes := validation.Attributes(map[string]string{
		"input":         "archivo de entrada",
		"output_format": "formato de salida",
		"output_dir":    "ruta de salida",
	})
	validator, err := facades.Validation().Make(data, rules, messages, attributes)
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"error":     err.Error() + ".",
			"message":   nil,
			"validator": nil,
		})
	} else if validator.Fails() {
		return ctx.Response().Status(400).Json(http.Json{
			"error":     nil,
			"message":   nil,
			"validator": validator.Errors().All(),
		})
	}

	cmd := exec.Command(
		facades.Storage().Disk("public").Path("whisper\\whisper-faster.exe"),
		[]string{
			fmt.Sprintf("--model=%s", "medium"),
			fmt.Sprintf("%s", data["input"]),
			fmt.Sprintf("--language=%s", "Spanish"),
			fmt.Sprintf("%s", "-pp"),
			fmt.Sprintf("--beam_size=%s", "1"),
			fmt.Sprintf("--best_of=%s", "1"),
			fmt.Sprintf("--output_format=%s", "all"),
			fmt.Sprintf("--output_dir=%s", data["output_dir"]),
		}...)

	// Create a new file for logging requests.
	logFile, err := os.Create(facades.Storage().Disk("public").Path("logs\\logs.log"))
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer, logFile)
	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		if err != nil {
			return ctx.Response().Success().Json(http.Json{
				"success": false,
				"message": "Se ha forzado la terminacion del subproceso.",
			})
		}
	}

	return ctx.Response().Success().Json(http.Json{
		"success": true,
		"message": "Procesamiento terminado con exito.",
	})
}

func (r *ConversionController) Stop(ctx http.Context) http.Response {
	cmd := exec.Command(
		"taskkill",
		[]string{
			fmt.Sprintf("%s", "/F"),
			fmt.Sprintf("%s", "/IM"),
			fmt.Sprintf("%s", "whisper-faster.exe"),
		}...)
	err := cmd.Run()
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"error": err,
		})
	}
	return ctx.Response().Success().Json(http.Json{
		"message": "Procesamiento cancelado.",
	})
}

func (r *ConversionController) RemoveLogs(ctx http.Context) http.Response {
	err := os.Remove(path.Storage() + "\\app\\public\\logs\\logs.log")
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"error": err,
		})
	}
	return ctx.Response().Success().Json(http.Json{
		"message": "Logs eliminados.",
	})
}
