package main

import (
	"github.com/ithaiq/tgin/example/classes"
	"github.com/ithaiq/tgin/example/middlewares"
	"github.com/ithaiq/tgin/internal"
)

func main() {
	internal.NewTGin().
		Attach(middlewares.NewUserMid()).
		Mount("v1", classes.NewIndexClass()).
		Launch()
}
