package service

import "github.com/bruhyan/linkpls/internal/database"

type LinkService struct {
	cache database.Cache
}

func New(cache database.Cache) *LinkService {
	return &LinkService{
		cache: cache,
	}
}

func (s *LinkService) CreateShortLink(url string) (string, error) {
	return "short_link", nil
}
