package handler

import (
	"github.com/LKarlon/test_task/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type service interface {
	GetInfo(serverID int) (res models.Servers, err error)
	DeleteInfo(serverID int) (er error)
}

type transport interface {
	ServerIdDecode(ctx *gin.Context) (id int, err error)
	//GetInfoEncode(ctx *gin.Context, res models.Servers) (err error)
}

type Handler struct {
	service   service
	transport transport
}

func NewHandler(services service, transport transport) *Handler {
	return &Handler{service: services, transport: transport}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/:id", h.GetInfo)
	router.DELETE("/:id", h.DeleteInfo)

	return router
}

func (h *Handler) GetInfo(ctx *gin.Context) {
	id, err := h.transport.ServerIdDecode(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	info, err := h.service.GetInfo(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, info)
}

func (h *Handler) DeleteInfo(ctx *gin.Context) {
	id, err := h.transport.ServerIdDecode(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.DeleteInfo(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
