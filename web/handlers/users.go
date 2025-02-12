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
	h.POST((""), h.post)
	h.PUT(("/:id"), h.put)
	h.DELETE(("/:id"), h.deleteUser)

}

// @Summary		ユーザ一覧取得
// @Description	ユーザ一覧を取得します
// @Accept			json
// @Produce		json
// @Success		200	{object}	models.UserIndexResponseModel
// @Router			/users [get]
// @Tags			users
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

// @Summary		ユーザ取得
// @Description	ユーザを取得します
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"ユーザID"
// @Success		200	{object}	models.UserGetResponseModel
// @Router			/users/{id} [get]
// @Tags			users
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

// @Summary		ユーザ登録
// @Description	ユーザを登録します
// @Accept			json
// @Produce		json
// @Param user body models.UserPostRequestModel true "ユーザー"
// @Success		200	{object}	models.UserPostResponseModel
// @Router			/users [post]
// @Tags			users
func (h *UserHandler) post(c echo.Context) error {
	req := new(models.UserPostRequestModel)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, map[string]string{
			"message": err.Error(),
		})
	}

	input := ports.NewUserRegisterInputData(req.Name)
	output, err := h.register.Handle(input)
	if err != nil {
		return c.JSON(500, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(200, models.UserPostResponseModel{
		CreatedId: output.CreatedUserID(),
	})
}

// @Summary		ユーザ更新
// @Description	ユーザを更新します
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"ユーザID"
// @Param user body models.UserPutRequestModel true "ユーザー"
// @Success		200	{object}	models.UserPostResponseModel
// @Router			/users/{id} [put]
// @Tags			users
func (h *UserHandler) put(c echo.Context) error {
	req := new(models.UserPutRequestModel)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, map[string]string{
			"message": err.Error(),
		})
	}

	input := ports.NewUserUpdateInputData(c.Param("id"), req.Name)
	_, err := h.update.Handle(input)
	if err != nil {
		return c.JSON(500, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(200, map[string]string{
		"message": "success",
	})
}

// @Summary		ユーザ削除
// @Description	ユーザを削除します
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"ユーザID"
// @Success		200	{object}	map[string]string
// @Router			/users/{id} [delete]
// @Tags			users
func (h *UserHandler) deleteUser(c echo.Context) error {
	input := ports.NewUserDeleteInputData(c.Param("id"))
	_, err := h.delete.Handle(input)
	if err != nil {
		return c.JSON(500, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(200, map[string]string{
		"message": "success",
	})
}
