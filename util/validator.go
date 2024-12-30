package util

import "github.com/go-playground/validator/v10"

func LevelNameValidator(fl validator.FieldLevel) bool {
	LevelName := fl.Field().String()
	allowedLevels := []string{"Circle", "Region", "City"}
	for _, Level := range allowedLevels {
		if LevelName == Level {
			return true
		}
	}
	return false
}
