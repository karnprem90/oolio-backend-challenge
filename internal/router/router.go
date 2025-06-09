package router

import (
	"oolio-backend-challenge/internal/handler"
	"oolio-backend-challenge/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func New() *Router {
	return &Router{
		engine: gin.Default(),
	}
}

func (r *Router) Public() *gin.RouterGroup {
	return r.engine.Group("/")
}

func (r *Router) Protected() *gin.RouterGroup {
	group := r.engine.Group("/")
	group.Use(middleware.AuthMiddleware())
	return group
}

func (r *Router) Engine() *gin.Engine {
	return r.engine
}

func SetupRouter(
	productHandler *handler.ProductHandler,
	orderHandler *handler.OrderHandler,
	promoHandler *handler.PromoHandler,
) *gin.Engine {
	router := gin.Default()

	public := router.Group("/")
	{
		public.GET("/product", productHandler.ListProducts)
		public.GET("/product/:id", productHandler.GetProduct)
	}

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/order", orderHandler.PlaceOrder)
		protected.GET("/order/:id", orderHandler.GetOrder)
		protected.GET("/orders", orderHandler.ListOrders)

		// Promo routes
		protected.GET("/validate-promo/:code", promoHandler.ValidatePromoCode)
	}

	return router
}
