package main

import (
	"fmt"

	"github.com/icyre/icyre/internal/config"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {

	cfg := config.New()

	fmt.Println(cfg)

}
