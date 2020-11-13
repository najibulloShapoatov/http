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

//Banner ..Структура нашего баннера
type Banner struct {
	ID      int64
	Title   string
	Content string
	Button  string
	Link    string
}

//это стартовый ID но для каждого создание баннера его изменяем
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
		//если ID элемента равно ID из параметра то мы нашли баннер
		if v.ID == id {
			//вернем баннер и ошибку nil
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
		//увеличиваем стартовый ID
		sID++
		//выставляем новый ID для баннера
		item.ID = sID
		//добавляем его в слайс
		s.items = append(s.items, item)
		//вернем баннер и ошибку nil
		return item, nil
	}
	//если id не равно 0 то ишем его из сушествуеших
	for k, v := range s.items {
		//если нашли то заменяем старый баннер с новым
		if v.ID == item.ID {
			//если нашли то в слайс под индексом найденного выставим новый элемент
			s.items[k] = item
			//вернем баннер и ошибку nil
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
			//берем все элементы до найденного и добавляем в него все элементы после найденного
			s.items = append(s.items[:k], s.items[k+1:]...)
			//вернем баннер и ошибку nil
			return v, nil
		}
	}

	//если не нашли то вернем ошибку что у нас такого банера не сушествует
	return nil, errors.New("item not found")
}
