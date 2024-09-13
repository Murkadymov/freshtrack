package freshtrack

type Repository interface {
	AddSupply() error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddSupply() {
	const op = "service.freshtrack.AddSupply"

	s.repo.AddSupply()

}
