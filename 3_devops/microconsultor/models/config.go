package models

type Config struct {
	Micro   string `validate:"required"`
	Env     string `validate:"required"`
	Version string `validate:"required"`
}
