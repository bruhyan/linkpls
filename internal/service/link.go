package service

type LinkService struct {
	// todo: dependencies
}

func (s *LinkService) CreateShortLink(url string) (string, error) {
	return "short_link", nil
}
