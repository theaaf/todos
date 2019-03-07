package app

import "github.com/sirupsen/logrus"

type App struct {
	Config *Config
}

func (a *App) NewContext() *Context {
	return &Context{
		Logger: logrus.New(),
	}
}

func New() (app *App, err error) {
	app = &App{}
	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
	return app, err
}
