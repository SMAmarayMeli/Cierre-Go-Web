package routes

import (
	"Cierre-Go-Web/cmd/handlers"
	"Cierre-Go-Web/internal/domain"
	tickets "Cierre-Go-Web/internal/product"
	"github.com/gin-gonic/gin"
)

type Router struct {
	db *[]domain.Ticket
	en *gin.Engine
}

func NewRouter(en *gin.Engine, db *[]domain.Ticket) *Router {
	return &Router{en: en, db: db}
}

func (r *Router) SetRoutes() {
	r.SetWebsite()
}

// website
func (r *Router) SetWebsite() {
	// instances
	rp := tickets.NewRepository(*r.db)
	sv := tickets.NewService(rp)
	h := handlers.NewTicket(sv)

	prod := r.en.Group("/ticket")

	prod.GET("/", h.GetAll())
	prod.GET("/getByCountry/:dest", h.GetTicketsByCountry())
	prod.GET("/getAverage/:dest", h.AverageDestination())

}
