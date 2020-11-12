package banners

import (
	"context"
	"errors"
	"sync"
)

//Service .  Это сервис для управления баннерами
type Service struct {
	mu    sync.RWMutex
	items []*Banner
}

//NewService . функция для создания нового сервиса
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

//это стартовый ID но для каждого создание поста его изменяем
var sID int64 = 0

//All ...
func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	//вернем все баннеры если их нет просто там окажется []
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

	//Проверяем если id равно 0 то создаем баннер
	if item.ID == 0 {
		//увеличиваем стартовый индекс
		sID++
		//выставляем новый индекс для баннера
		item.ID = sID
		//добавляем его в слайс
		s.items = append(s.items, item)
		return item, nil
	}
	//если нет то ишем его из сушествуеших
	for k, v := range s.items {
		//если нашли то заменяем старый баннер с новым
		if v.ID == item.ID {
			s.items[k] = item
			return item, nil
		}
	}
	//если не нашли то вернем ошибку что у нас такого банера не сушествует 
	return nil, errors.New("item not found")
}

//RemoveByID ... Метод для удаления 
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	//ишем баннер из слайса
	for k, v := range s.items {
		//если нашли то удаляем его из слайса
		if v.ID == id {
			s.items = removeIndex(s.items, k)
			return v, nil
		}
	}

	//если не нашли то вернем ошибку что у нас такого банера не сушествует 
	return nil, errors.New("item not found")
}

//Функция который удаляет элемент из слайса
func removeIndex(s []*Banner, index int) []*Banner {
	return append(s[:index], s[index+1:]...)
}


