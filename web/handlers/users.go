package handlers

import (
	"go-arch-practice/usecase/users/ports"
	"go-arch-practice/web/models"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	*echo.Group
	get      ports.IUserGetInputPort
	get_all  ports.IUserGetAllInputPort
	update   ports.IUserUpdateInputPort
	register ports.IUserRegisterInputPort
	delete   ports.IUserDeleteInputPort
}

func NewUserHandler(
	g *echo.Group,
	get ports.IUserGetInputPort,
	get_all ports.IUserGetAllInputPort,
	update ports.IUserUpdateInputPort,
	register ports.IUserRegisterInputPort,
	delete ports.IUserDeleteInputPort,
) *UserHandler {
	return &UserHandler{
		Group:    g,
		get:      get,
		get_all:  get_all,
		update:   update,
		register: register,
		delete:   delete,
	}
}

func (h *UserHandler) Regist() {

	h.GET((""), h.index)
	h.GET(("/:id"), h.getByID)

}

func (h *UserHandler) index(c echo.Context) error {
	users, err := h.get_all.Handle()
	if err != nil {
		return c.JSON(500, map[string]string{
			"message": err.Error(),
		})
	}

	responseUsers := make([]models.UserResponseModel, len(users.Users()))
	for i, user := range users.Users() {
		responseUsers[i] = models.UserResponseModel{
			Id:   user.ID(),
			Name: user.Name(),
		}
	}

	return c.JSON(200, models.UserIndexResponseModel{
		Users: responseUsers,
	})
}

func (h *UserHandler) getByID(c echo.Context) error {
	input := ports.NewUserGetInputData(c.Param("id"))
	output, err := h.get.Handle(input)
	if err != nil {
		return c.JSON(500, map[string]string{
			"message": err.Error(),
		})
	}

	userData := (*output).User()
	return c.JSON(200, models.UserGetResponseModel{
		UserResponseModel: models.UserResponseModel{
			Id:   userData.ID(),
			Name: userData.Name(),
		},
	})
}
