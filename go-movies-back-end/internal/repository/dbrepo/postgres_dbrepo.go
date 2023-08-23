package dbrepo

import (
	"backend/internal/models"
	"database/sql"
	"time"
	"context"
)

type postgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *postgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `
		select
			id, title, release_date, runtime, mpaa_rating, description, coalesce(image, ''), created_at, updated_at

		from
			movies
		order by
			title
		`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}