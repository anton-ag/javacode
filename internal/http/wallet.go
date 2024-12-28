package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payload struct {
	ID        string `json:"walletId"`
	Operation string `json:"operationType"`
	Amount    int    `json:"amount"`
}

type Balance struct {
	Amount int `json:"amount"`
}

func (h *Handler) initWalletRoutes(api *gin.RouterGroup) {
	wallet := api.Group("/wallet")
	{
		wallet.POST("/", h.Update)
		wallet.GET("/:id", h.Check)
	}
}

func (h *Handler) Update(c *gin.Context) {
	var p Payload
	if err := c.BindJSON(&p); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Wallet.Update(p.ID, p.Amount, p.Operation); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Check(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newResponse(c, http.StatusBadRequest, "Не указан ID кошелька")
		return
	}

	amount, err := h.service.Wallet.Check(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Balance{amount})
}
