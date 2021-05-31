package handler

import (
	"io/ioutil"

	"github.com/LKarlon/test_task/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/metrics", h.convert)

	return router
}

func (h *Handler) convert(c *gin.Context){	 
	file, err := ioutil.ReadFile("./file.yaml")
	if err != nil {
		logrus.Errorf("file read error: %s", err.Error())
		return
	}

	str, err := h.services.Convert(file)
	if err != nil {
		logrus.Errorf("service error: %s", err.Error())
		return
	}
	c.String(200, str)
}