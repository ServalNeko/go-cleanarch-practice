package main

import (
	"go-arch-practice/infra/pg"
	user_repo "go-arch-practice/infra/pg/persistence/users"
	user_usecase "go-arch-practice/usecase/users/interactor"
	"go-arch-practice/web/handlers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "go-arch-practice/cmd/server/docs"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8000
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

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
