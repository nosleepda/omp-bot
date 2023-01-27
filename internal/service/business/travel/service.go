package travel

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"github.com/samber/lo"
	"time"
)

type TravelService interface {
	Describe(travelId int64) (*business.Travel, error)
	List(cursor int64, limit int64) ([]*business.Travel, error)
	Count() (int64, error)
	Create(business.Travel) (*business.Travel, error)
	Update(travelId int64, travel business.Travel) error
	Remove(travelId int64) (bool, error)
}

type DummyTravelService struct {
	Data []*business.Travel
}

func NewDummyTravelService() *DummyTravelService {
	dummyData := getDummyData()

	return &DummyTravelService{
		Data: dummyData,
	}
}

func (s *DummyTravelService) Count() (int64, error) {
	return int64(len(s.Data)), nil
}

func (s *DummyTravelService) Describe(travelId int64) (*business.Travel, error) {
	travels := lo.Filter(s.Data, func(travel *business.Travel, index int) bool {
		return travel.TravelId == travelId
	})

	if len(travels) == 1 {
		return travels[0], nil
	}

	err := fmt.Errorf("travel not found by %v", travelId)

	return nil, err
}

func (s *DummyTravelService) List(cursor int64, limit int64) ([]*business.Travel, error) {

	if cursor < 0 {
		return nil, errors.New("cursor must be greater or equal to 0")
	}

	if limit < 0 {
		return nil, errors.New("limit must be greater to 0")
	}

	count := int64(len(s.Data))

	if cursor >= count {
		return nil, errors.New("no data for show")
	}

	if cursor+limit >= count {
		return s.Data[cursor:], nil
	}

	return s.Data[cursor : cursor+limit], nil
}

func (s *DummyTravelService) Create(travel business.Travel) (*business.Travel, error) {

	ids := lo.Map(s.Data, func(travel *business.Travel, index int) int64 {
		return travel.TravelId
	})

	maxId := lo.Max(ids)

	travel.TravelId = maxId + 1

	s.Data = append(s.Data, &travel)

	return &travel, nil
}

func (s *DummyTravelService) Update(travelId int64, travelUpsert business.Travel) error {
	travels := lo.Filter(s.Data, func(travel *business.Travel, index int) bool {
		return travel.TravelId == travelId
	})

	if len(travels) != 1 {

		return fmt.Errorf("travel not found by %v", travelId)
	}
	travel := travels[0]

	travel.Where = travelUpsert.Where
	travel.Title = travelUpsert.Title
	travel.StartDate = travelUpsert.StartDate
	travel.Duration = travelUpsert.Duration

	return nil
}

func (s *DummyTravelService) Remove(travelId int64) (bool, error) {
	travels := lo.Filter(s.Data, func(travel *business.Travel, index int) bool {
		return travel.TravelId == travelId
	})

	if len(travels) != 1 {
		return false, fmt.Errorf("travel not found by %v", travelId)
	}

	s.Data = lo.Filter(s.Data, func(travel *business.Travel, index int) bool {
		return travel.TravelId != travelId
	})

	return true, nil
}

func getDummyData() []*business.Travel {
	return []*business.Travel{
		{
			Title:     "Travel in Turkey",
			Where:     "Istanbul",
			StartDate: time.Date(2022, 9, 7, 0, 0, 0, 0, time.Local),
			Duration:  14,
		},
		{
			Title:     "Travel in Germany",
			Where:     "Berlin",
			StartDate: time.Date(2022, 12, 7, 0, 0, 0, 0, time.Local),
			Duration:  10,
		},
		{
			Title:     "Travel in Turkey",
			Where:     "Antalya",
			StartDate: time.Date(2023, 1, 3, 0, 0, 0, 0, time.Local),
			Duration:  4,
		},
		{
			Title:     "Travel in Georgia",
			Where:     "Tbilisi",
			StartDate: time.Date(2023, 2, 14, 0, 0, 0, 0, time.Local),
			Duration:  10,
		},
		{
			Title:     "Travel in Canada",
			Where:     "Toronto",
			StartDate: time.Date(2023, 3, 7, 0, 0, 0, 0, time.Local),
			Duration:  7,
		},
	}
}
