package wrap

import (
	"time"
)

type Service struct {
	repo wrapinterface
}

func NewService(repo wrapinterface) ServiceInterface {
	return &Service{
		repo: repo,
	}
}
func (s *Service) MakeWrap(name string) (*Wrap, error) {
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

func (s *Service) GetAllWraps() ([]*Wrap, error) {
	wraps, err := s.repo.GetAllWraps()
	if err != nil {
		return nil, err
	}
	return wraps, nil
}
