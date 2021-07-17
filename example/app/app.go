package app

import (
	"github.com/ravielze/oculi/app"
	"github.com/ravielze/oculi/example/di"
	"github.com/ravielze/oculi/example/infrastructures"
	"github.com/ravielze/oculi/example/resources"

	//webserver "github.com/ravielze/oculi/server/gin"
	"go.uber.org/dig"
)

func Run() {
	invoker := func(container *dig.Container) error {
		return container.Invoke(func(i infrastructures.Component, r resources.Resource) error {
			//return webserver.New(i, r).Run()
			return nil
		})
	}

	if err := app.Run(di.Container, invoker); err != nil {
		panic(err)
	}
}
