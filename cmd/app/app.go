package main

import (
	"github.com/icyre/icyre/internal/app"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {

	a, err := app.New()
	if err != nil {
		panic(err)
	}

	a.Run()

}
