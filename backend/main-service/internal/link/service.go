package link

import (
	"fmt"
	"main-service/internal/models"
)

// Service provides business logic for link operations
type Service struct {
	repo *Repository
}

// NewService creates new Service instance
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Create generates short link for original URL
func (s *Service) Create(originLink string) (string, error) {
	if originLink == "" {
		return "", fmt.Errorf("error: Empty link")
	}
	exists, err := s.repo.Exist(originLink)
	if err != nil {
		return "", fmt.Errorf("error checking link existence: %v", err)
	}
	if exists {
		return "", fmt.Errorf("error: Link already exists")
	}
	pseudoLink := generatePseudoLink()
	return pseudoLink, s.repo.Create(&models.Link{
		OriginLink: originLink,
		PseudoLink: pseudoLink,
	})
}

// GetLink retrieves link by short URL
func (s *Service) GetLink(link string) (*models.Link, error) {
	if link == "" {
		return nil, fmt.Errorf("error: Empty link")
	}
	linkObject, err := s.repo.Find(link)
	return linkObject, err
}

// DeleteLink removes link by original URL
func (s *Service) DeleteLink(originLink string) error {
	if originLink == "" {
		return fmt.Errorf("error: Empty link")
	}
	return s.repo.Delete(originLink)

}
