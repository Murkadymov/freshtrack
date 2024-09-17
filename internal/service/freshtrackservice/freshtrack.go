package freshtrackservice

import (
	"fmt"
	"freshtrack/internal/entity"
	"log"
	"log/slog"
)

type Repository interface {
	AddSupply(supply *entity.Supply) error
	GetSupplyList() ([]entity.Supply, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddSupply(e *entity.Supply) error {
	const op = "service.freshtrackservice.AddSupply"

	err := s.repo.AddSupply(e)
	if err != nil {
		slog.Error(op+"addSupply", slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) GetSupplyList() ([]entity.Supply, error) {
	const op = "service.freshtrackservice."
	log.Println("started service getlist")
	supplyList, err := s.repo.GetSupplyList()
	if err != nil {
		return nil, fmt.Errorf("%s + GetSupplyList: %w", op, err)
	}
	log.Println("ended service getlist")

	return supplyList, nil
}
