package router

import (
	"oolio-backend-challenge/internal/handler"
)

func (r *Router) RegisterProductRoutes(handler *handler.ProductHandler) {
	public := r.Public()
	{
		public.GET("/product", handler.ListProducts)
		public.GET("/product/:id", handler.GetProduct)
	}
}
