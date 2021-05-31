package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Transport interface {
	ServerIdDecode(ctx *gin.Context) (id int, err error)
}

type transport struct {
}

func (t *transport) ServerIdDecode(ctx *gin.Context) (id int, err error) {
 	res := ctx.Param("id")
	if res == "" {
		return
	}
	id, err = strconv.Atoi(res)
	return
}


func NewTransport() Transport {
	return &transport{}
}