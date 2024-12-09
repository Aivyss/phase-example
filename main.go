package main

import (
	"github.com/labstack/echo/v4"
	"phase_example/handler"
	"phase_example/usecase"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	e := echo.New()

	accountHandler := handler.NewAccountHandler(usecase.NewAccountUsecase())
	e.POST("/account/signup", accountHandler.PostAccountSignup)

	if err := e.Start(":1234"); err != nil {
		panic(err)
	}
}
