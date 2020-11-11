package banners

import (
	"context"
	"errors"
	"sync"
)

//Service ...
type Service struct {
	mu    sync.RWMutex
	items []*Banner
}

//NewService ...
func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}
}

//Banner ..
type Banner struct {
	ID      int64
	Title   string
	Content string
	Button  string
	Link    string
}

var sID int64 = 0

//All ...
func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items, nil
}

//ByID ...
func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range s.items {
		if v.ID == id {
			return v, nil
		}
	}

	return nil, errors.New("item not found")
}

//Save ...
func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if item.ID == 0 {
		sID = sID + 1
		item.ID = sID
		s.items = append(s.items, item)
		return item, nil
	}
	for k, v := range s.items {
		if v.ID == item.ID {
			s.items[k] = item
			return item, nil
		}
	}

	return nil, errors.New("item not found")
}

//RemoveByID ...
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, v := range s.items {
		if v.ID == id {
			s.items = removeIndex(s.items, k)
			return v, nil
		}
	}

	return nil, errors.New("item not found")
}

func removeIndex(s []*Banner, index int) []*Banner {
	return append(s[:index], s[index+1:]...)
}
