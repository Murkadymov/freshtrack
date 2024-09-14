package freshtrackservice

import (
	"fmt"
	"freshtrack/internal/entity"
)

type Repository interface {
	AddSupply(supply entity.Supply) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddSupply(e entity.Supply) error {
	const op = "service.freshtrackrepo.AddSupply"

	err := s.repo.AddSupply(e)
	if err != nil {
		fmt.Println("got error")
		return err
	}

	return nil
}
