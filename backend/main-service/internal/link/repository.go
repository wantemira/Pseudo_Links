package link

import (
	"context"
	"fmt"
	"main-service/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
)

// Repository handles link data persistence
type Repository struct {
	client *redis.Client
}

// NewRepository creates new Repository instance
func NewRepository(client *redis.Client) *Repository {
	return &Repository{client: client}
}

// Create saves new link to database
func (r *Repository) Create(link *models.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.client.Set(ctx, link.OriginLink, link.PseudoLink, 0).Err(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := r.client.Set(ctx, link.PseudoLink, link.OriginLink, 0).Err(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

// Find retrieves link by original URL
func (r *Repository) Find(originLink string) (link *models.Link, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pseudoLink, err := r.client.Get(ctx, originLink).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // key not found
		}
		return nil, err
	}

	return &models.Link{
		OriginLink: originLink,
		PseudoLink: pseudoLink,
	}, nil
}

// Delete removes link from database
func (r *Repository) Delete(originLink string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pseudoLink, err := r.client.Get(ctx, originLink).Result()
	if err != nil {
		if err != redis.Nil {
			return err
		}
	}
	if err := r.client.Del(ctx, originLink).Err(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := r.client.Del(ctx, pseudoLink).Err(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil

}

// Exist checks if link exists in database
func (r *Repository) Exist(originLink string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exists, err := r.client.Exists(ctx, originLink).Result()
	return exists > 0, err
}
