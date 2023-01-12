package tickets

import (
	"Cierre-Go-Web/internal/domain"
	"context"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

func NewService(rp Repository) Repository {
	return &service{rp: rp}
}

type service struct {
	rp Repository
}

func (sv *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	return sv.rp.GetAll(ctx)
}

func (sv *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

}
