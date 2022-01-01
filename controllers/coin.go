package controllers

import (
	"net/http"

	"github.com/WebDev-Naveen/go-backend/config"

	"github.com/gin-gonic/gin"
)

//OrderHandler --> Handler for Order Entity
type CoinHandler interface {
	UserCoin(*gin.Context)
}

type coinHandler struct {
	repo config.CoinRepository
}

//NewOrderHandler --> return new Order Handler
func NewCoinHandler() CoinHandler {
	return &coinHandler{
		repo: config.NewCoinRepository(),
	}
}

func (h *coinHandler) UserCoin(ctx *gin.Context) {

	userID := ctx.GetFloat64("userID")
	if err := h.repo.UserCoin(int(userID)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.String(http.StatusOK, "Product Successfully ordered")
	}
}
