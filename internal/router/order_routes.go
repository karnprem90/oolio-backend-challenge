package router

import (
	"oolio-backend-challenge/internal/handler"
)

func (r *Router) RegisterOrderRoutes(handler *handler.OrderHandler) {
	protected := r.Protected()
	{
		protected.POST("/order", handler.PlaceOrder)
		protected.GET("/order/:id", handler.GetOrder)
		protected.GET("/orders", handler.ListOrders)
	}
}
