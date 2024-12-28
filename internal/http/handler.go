package http

import (
	"github.com/anton-ag/javacode/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		h.initWalletRoutes(api)
	}
}
