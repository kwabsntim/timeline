package wrap

import (
	"time"
)

type Service struct {
	repo wrapinterface
}

func NewService(repo wrapinterface) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) MakeWrap(name, status string) (*Wrap, error) {
	wrap := &Wrap{
		Name:       name,
		Status:     StatusPending,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	err := s.repo.CreateWrap(wrap)
	if err != nil {
		return nil, err
	}
	return wrap, nil

}
func (s *Service) GetWrap(uuid string) (*Wrap, error) {
	wrap, err := s.repo.Getwrap(uuid)
	if err != nil {
		return nil, err
	}
	return wrap, nil
}
