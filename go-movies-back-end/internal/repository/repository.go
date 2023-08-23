package repository

import "backend/internal/models"

type Database interface {
	AllMovies() ([]*models.Movie, error)
}