package main

import (
	"Cierre-Go-Web/cmd/routes"
	"Cierre-Go-Web/internal/domain"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {

	// Cargo csv.
	list, err := LoadTicketsFromFile("internal/product/tickets.csv")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	router := routes.NewRouter(r, &list)
	router.SetRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
