package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/icyre/icyre/internal/config"
)

type App struct {
	cfg *config.Config
	app *gin.Engine
}

func newApp(cfg *config.Config) *App {
	app := gin.Default()
	return &App{cfg, app}
}

func (a *App) Run() error {

	return a.app.Run(fmt.Sprintf("%s:%d", a.cfg.App.Host, a.cfg.App.Port))
}
