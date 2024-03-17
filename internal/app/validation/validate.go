package validation

import "github.com/go-playground/validator/v10"

type App struct {
	Validate *validator.Validate
}

func New() *App {
	return &App{
		Validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
