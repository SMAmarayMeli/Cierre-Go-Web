package handlers

import (
	"Cierre-Go-Web/internal/product"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	service product.Service
}

func NewTicket(sv product.Service) *Service {
	return &Service{service: sv}
}

func (s *Service) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets, err := s.service.GetAll(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
