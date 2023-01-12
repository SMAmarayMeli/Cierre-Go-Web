package product

import (
	"Cierre-Go-Web/internal/domain"
	"context"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

func NewService(rp Repository) Service {
	return &service{rp: rp}
}

type service struct {
	rp Repository
}

func (sv *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	return sv.rp.GetAll(ctx)
}

func (sv *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	return sv.rp.GetTicketByDestination(ctx, destination)
}

func (sv *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	return sv.rp.GetTicketByDestination(ctx, destination)
}

func (sv *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	total, errTotal := sv.rp.GetAll(ctx)
	if errTotal != nil {
		return 0, errTotal
	}
	totalDest, errDest := sv.rp.GetTicketByDestination(ctx, destination)
	if errDest != nil {
		return 0, errDest
	}

	return float64(len(totalDest) * len(total) / 100), nil
}
