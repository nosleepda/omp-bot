package travel

import (
	"github.com/ozonmp/omp-bot/internal/model/business"
	"strconv"
	"time"
)

func mapTravel(fields []string) (*business.Travel, error) {
	title := fields[0]
	where := fields[1]
	date, err := time.Parse("1/29/2023", fields[2])
	if err != nil {
		return nil, err
	}
	duration, err := strconv.Atoi(fields[3])
	if err != nil {
		return nil, err
	}

	return &business.Travel{
		Title:     title,
		Where:     where,
		StartDate: date,
		Duration:  duration,
	}, nil
}
