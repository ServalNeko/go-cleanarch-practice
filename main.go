package main

import (
	"go-arch-practice/infra/pg"
	user_repo "go-arch-practice/infra/pg/persistence/users"
	user_usecase "go-arch-practice/usecase/users/interactor"
	"go-arch-practice/web/handlers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	var e = echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	db, err := pg.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	eg := e.Group("/users")
	uh := handlers.NewUserHandler(
		eg,
		user_usecase.NewUserGetInteractor(user_repo.NewUserRepository(db)),
		user_usecase.NewUserGetAllInteractor(user_repo.NewUserRepository(db)),
		user_usecase.NewUserUpdateInteractor(user_repo.NewUserRepository(db)),
		user_usecase.NewUserRegisterInteractor(user_repo.NewUserRepository(db)),
		user_usecase.NewUserDeleteInteractor(user_repo.NewUserRepository(db)),
	)
	uh.Regist()

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}

}
