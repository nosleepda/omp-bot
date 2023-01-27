package business

import (
	"fmt"
	"time"
)

type Travel struct {
	TravelId  int64
	Title     string
	Where     string
	StartDate time.Time
	Duration  int
}

func NewTravel(title string, where string, startDate time.Time, duration int) *Travel {
	return &Travel{
		Title:     title,
		Where:     where,
		StartDate: startDate,
		Duration:  duration,
	}
}

func (d *Travel) String() string {
	return fmt.Sprintf("%s\n%s\n%v\n", d.Title, d.Where, d.StartDate)
}
