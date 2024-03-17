package validation

import (
	"github.com/go-playground/validator/v10"
)

type App struct {
	Validate *validator.Validate
}

func New() *App {
	vlt := validator.New(validator.WithRequiredStructEnabled())

	return &App{
		Validate: vlt,
	}
}
