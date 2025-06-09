package handler

import (
	"errors"
	"net/http"

	"oolio-backend-challenge/internal/domain"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service domain.OrderService
}

func NewOrderHandler(service domain.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var req struct {
		Items []struct {
			ProductID string `json:"productId" binding:"required"`
			Quantity  int    `json:"quantity" binding:"required,min=1"`
		} `json:"items" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	items := make([]domain.OrderItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = domain.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	order, err := h.service.PlaceOrder(c.Request.Context(), items)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrProductNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		case errors.Is(err, domain.ErrInvalidQuantity):
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quantity"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to place order"})
		}
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrder handles GET /order/:id
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order ID is required"})
		return
	}

	order, err := h.service.GetOrder(c.Request.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrOrderNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get order"})
		}
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.service.ListOrders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
