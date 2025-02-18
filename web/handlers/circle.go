package handlers

import (
	"go-arch-practice/usecase/circles/ports"
	circle_ports "go-arch-practice/usecase/circles/ports"
	user_ports "go-arch-practice/usecase/users/ports"
	"go-arch-practice/web/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CircleHandler struct {
	*echo.Group
	get      circle_ports.ICircleGetInputPort
	get_all  circle_ports.ICircleGetAllInputPort
	register circle_ports.ICircleRegisterInputPort
	update   circle_ports.ICircleUpdateInputPort
	join     circle_ports.ICircleJoinInputPort
	delete   circle_ports.ICircleDeleteInputPort
	user_get user_ports.IUserGetInputPort
}

func NewCircleHandler(
	g *echo.Group,
	get circle_ports.ICircleGetInputPort,
	get_all circle_ports.ICircleGetAllInputPort,
	register circle_ports.ICircleRegisterInputPort,
	update circle_ports.ICircleUpdateInputPort,
	join circle_ports.ICircleJoinInputPort,
	delete circle_ports.ICircleDeleteInputPort,
	user_get user_ports.IUserGetInputPort,
) *CircleHandler {
	return &CircleHandler{
		Group:    g,
		get:      get,
		get_all:  get_all,
		register: register,
		update:   update,
		join:     join,
		delete:   delete,
		user_get: user_get,
	}
}

func (h *CircleHandler) Regist() {
	h.GET("", h.index)
	h.POST("", h.post)
}

// @Summary		サークル一覧取得
// @Description	サークル一覧を取得します
// @Accept			json
// @Produce		json
// @Success		200	{object}	models.CircleIndexResponseModel
// @Router			/circles [get]
// @Tags			circles
func (h *CircleHandler) index(c echo.Context) error {
	outputData, err := h.get_all.Handle()
	if err != nil {
		return err
	}

	circles := outputData.Circles()
	circleModels := make([]models.CircleResponseModel, len(circles))
	for i, c := range circles {
		circleModels[i] = models.CircleResponseModel{
			Id:        c.ID(),
			Name:      c.Name(),
			MemberIds: c.MemberIDs(),
		}
	}

	response := models.CircleIndexResponseModel{
		Circles: circleModels,
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary		サークル登録
// @Description	サークルを登録します
// @Accept			json
// @Produce		json
// @Param circle body models.CirclePostRequestModel true "サークル"
// @Success		200	{object}	models.CirclePostResponseModel
// @Router			/circles [post]
// @Tags			circles
func (h *CircleHandler) post(c echo.Context) error {
	req := &models.CirclePostRequestModel{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"message": err.Error()})
	}

	input := ports.NewCircleRegisterInputData(req.CircleName, req.OwnerID)
	output, err := h.register.Handle(input)
	if err != nil {
		return err
	}

	response := models.CirclePostResponseModel{
		CreatedId: output.CreatedCircleID(),
	}

	return c.JSON(http.StatusOK, response)
}
