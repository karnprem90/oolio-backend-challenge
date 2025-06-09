package router

import (
	"oolio-backend-challenge/internal/handler"
)

func (r *Router) RegisterPromoRoutes(handler *handler.PromoHandler) {
	protected := r.Protected()
	{
		protected.GET("/validate-promo/:code", handler.ValidatePromoCode)
	}
}
