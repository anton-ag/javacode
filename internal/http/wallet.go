package http

import (
	"database/sql"
	"net/http"

	"github.com/anton-ag/javacode/internal/domain"
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
	}
	wallets := api.Group("/wallets")
	{
		wallets.GET("/:id", h.Check)
	}
}

func (h *Handler) Update(c *gin.Context) {
	var p Payload
	if err := c.BindJSON(&p); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Wallet.Update(p.ID, p.Amount, p.Operation); err != nil {
		if err == domain.ErrWalletNotFound {
			newResponse(c, http.StatusNotFound, err.Error())
			return
		}
		if err == domain.ErrWrongOperation {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}
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
		if err == sql.ErrNoRows {
			newResponse(c, http.StatusNotFound, "кошелёк с данным uuid не найден")
			return
		}
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, Balance{amount})
}
