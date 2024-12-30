package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type KabKotaRepository struct {
	db *sql.DB
}

func NewKabKotaRepository(db *sql.DB) *KabKotaRepository {
	return &KabKotaRepository{db: db}
}

func (r *KabKotaRepository) GetAllSpatial(ctx context.Context) ([]KabKota, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT gid, city, region, circle, ST_AsGeoJSON(geom) as geom FROM public.kab_kota`)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var result []KabKota
	for rows.Next() {
		var kabKota KabKota
		if err := rows.Scan(&kabKota.GID, &kabKota.City, &kabKota.Region, &kabKota.Circle, &kabKota.Geom); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		result = append(result, kabKota)
	}

	if len(result) == 0 {
		return nil, errors.New("no records found in kab_kota table")
	}

	return result, nil
}

type GetSpatialLv struct {
	Level     string `json:"level"`
	LevelName string `json:"level_name"`
}

func (r *KabKotaRepository) GetSpatialByFilter(ctx context.Context, arg GetSpatialLv) ([]KabKota, error) {
	// Validate input
	if arg.Level == "" || arg.LevelName == "" {
		return nil, errors.New("invalid filter arguments: level and level_name cannot be empty")
	}

	query := fmt.Sprintf(
		`SELECT gid, city, region, circle, ST_AsGeoJSON(geom) as geom FROM public.kab_kota WHERE %s = $1`,
		arg.Level,
	)

	rows, err := r.db.QueryContext(ctx, query, arg.LevelName)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query with filter %s=%s: %w", arg.Level, arg.LevelName, err)
	}
	defer rows.Close()

	var result []KabKota
	for rows.Next() {
		var kabKota KabKota
		if err := rows.Scan(&kabKota.GID, &kabKota.City, &kabKota.Region, &kabKota.Circle, &kabKota.Geom); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		result = append(result, kabKota)
	}

	// Check if no rows were returned
	if len(result) == 0 {
		return nil, fmt.Errorf("no records found for filter %s = %s", arg.Level, arg.LevelName)
	}

	return result, nil
}
