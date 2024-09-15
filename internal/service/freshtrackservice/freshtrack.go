package freshtrackservice

import (
	"fmt"
	"freshtrack/internal/entity"
	"log"
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
	const op = "service.freshtrackrepo.AddSupply"

	err := s.repo.AddSupply(e)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) GetSupplyList() ([]entity.Supply, error) {
	log.Println("started service getlist")
	supplyList, _ := s.repo.GetSupplyList()
	log.Println("ended service getlist")
	return supplyList, nil

}
