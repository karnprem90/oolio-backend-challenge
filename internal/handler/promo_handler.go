package handler

import (
	"net/http"

	"oolio-backend-challenge/internal/domain"

	"github.com/gin-gonic/gin"
)

type PromoHandler struct {
	service domain.PromoCodeService
}

func NewPromoHandler(service domain.PromoCodeService) *PromoHandler {
	return &PromoHandler{service: service}
}

func (h *PromoHandler) ValidatePromoCode(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "promo code is required"})
		return
	}

	valid, err := h.service.ValidatePromoCode(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": valid})
}
