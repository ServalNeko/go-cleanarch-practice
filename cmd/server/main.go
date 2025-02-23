package main

import (
	"go-arch-practice/domain/circles"
	circle_domain "go-arch-practice/domain/circles"
	"go-arch-practice/infra/pg"
	"go-arch-practice/infra/pg/persistence/repo"
	circle_usecase "go-arch-practice/usecase/circles/interactor"
	user_usecase "go-arch-practice/usecase/users/interactor"
	"go-arch-practice/web/handlers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	eg := e.Group("/users")
	uh := handlers.NewUserHandler(
		eg,
		user_usecase.NewUserGetInteractor(repo.NewUserRepository(db)),
		user_usecase.NewUserGetAllInteractor(repo.NewUserRepository(db)),
		user_usecase.NewUserUpdateInteractor(repo.NewUserRepository(db)),
		user_usecase.NewUserRegisterInteractor(repo.NewUserRepository(db)),
		user_usecase.NewUserDeleteInteractor(repo.NewUserRepository(db)),
	)
	uh.Regist()

	ec := e.Group("/circles")
	ch := handlers.NewCircleHandler(
		ec,
		circle_usecase.NewCircleGetInteractor(repo.NewCircleRepository(db)),
		circle_usecase.NewCircleGetAllInteractor(repo.NewCircleRepository(db)),
		circle_usecase.NewCircleRegisterInteractor(repo.NewCircleRepository(db), repo.NewUserRepository(db), circles.NewCircleService(repo.NewCircleRepository(db))),
		circle_usecase.NewCircleUpdateInteractor(repo.NewCircleRepository(db), circle_domain.NewCircleService(repo.NewCircleRepository(db))),
		circle_usecase.NewCircleJoinInteractor(repo.NewCircleRepository(db), repo.NewUserRepository(db)),
		circle_usecase.NewCircleDeleteInteractor(repo.NewCircleRepository(db)),
		user_usecase.NewUserGetInteractor(repo.NewUserRepository(db)),
	)

	ch.Regist()

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}

}
