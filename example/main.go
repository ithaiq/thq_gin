package main

import (
	"fmt"
	"github.com/ithaiq/tgin/example/classes"
	"github.com/ithaiq/tgin/example/middlewares"
	"github.com/ithaiq/tgin/internal"
)

func main() {
	internal.NewTGin().
		Beans(internal.NewGormAdapter()).
		Attach(middlewares.NewUserMid()).
		Mount("v1", classes.NewIndexClass()).
		Task("0/3 * * * * *", func() {
			fmt.Println("定时执行")
		}).
		Launch()
}
