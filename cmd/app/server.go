package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/najibulloShapoatov/http/pkg/banners"
)

//Server .. это наш логический сервер
type Server struct {
	mux       *http.ServeMux
	bannerSvc *banners.Service
}

//NewServer .. Функция для создание нового сервера
func NewServer(m *http.ServeMux, bnrSvc *banners.Service) *Server {
	return &Server{mux: m, bannerSvc: bnrSvc}
}

//ServeHTTP ... метод для запуска сервера
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

//Init .. мотод для инициализации сервера
func (s *Server) Init() {
	//здес мы зеристрируем роутеры с определенными хендлерами
	s.mux.HandleFunc("/banners.getAll", s.handleGetAllBanners)
	s.mux.HandleFunc("/banners.getById", s.handleGetBannerByID)
	s.mux.HandleFunc("/banners.save", s.handleSaveBanner)
	s.mux.HandleFunc("/banners.removeById", s.handleRemoveByID)
}

// хендлер метод для извлечения всех баннеров
func (s *Server) handleGetAllBanners(w http.ResponseWriter, r *http.Request) {

	//берем все баннеры из сервиса
	banners, err := s.bannerSvc.All(r.Context())

	//если получили какую нибуд ошибку то отвечаем с ошибкой
	if err != nil {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusInternalServerError)
		return
	}

	//преобразуем данные в JSON
	data, err := json.Marshal(banners)

	//если получили ошибку то отвечаем с ошибкой
	if err != nil {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusInternalServerError)
		return
	}

	//вызываем функцию для ответа в формате JSON
	respondJSON(w, data)

}

func (s *Server) handleGetBannerByID(w http.ResponseWriter, r *http.Request) {
	//получаем ID из параметра запроса
	idP := r.URL.Query().Get("id")

	// переобразуем его в число
	id, err := strconv.ParseInt(idP, 10, 64)
	//если получили ошибку то отвечаем с ошибкой
	if err != nil {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusBadRequest)
		return
	}

	//получаем баннер из сервиса
	banner, err := s.bannerSvc.ByID(r.Context(), id)
	//если получили ошибку то отвечаем с ошибкой
	if err != nil {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusInternalServerError)
		return
	}

	//преобразуем данные в JSON
	data, err := json.Marshal(banner)

	//если получили ошибку то отвечаем с ошибкой
	if err != nil {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusInternalServerError)
		return
	}

	//вызываем функцию для ответа в формате JSON
	respondJSON(w, data)
}

func (s *Server) handleSaveBanner(w http.ResponseWriter, r *http.Request) {

	//получаем данные из параметра запроса
	idP := r.URL.Query().Get("id")
	title := r.URL.Query().Get("title")
	content := r.URL.Query().Get("content")
	button := r.URL.Query().Get("button")
	link := r.URL.Query().Get("link")

	id, err := strconv.ParseInt(idP, 10, 64)
	//если получили ошибку то отвечаем с ошибкой
	if err != nil {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusBadRequest)
		return
	}
	//Здесь опционалная проверка то что если все данные приходит пустыми то вернем ошибку
	if title == "" && content == "" && button == "" && link == "" {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusBadRequest)
		return
	}

	//создаём указател на структуру баннера
	item := &banners.Banner{
		ID:      id,
		Title:   title,
		Content: content,
		Button:  button,
		Link:    link,
	}

	//вызываем метод Save тоест сохраняем или обновляем его
	banner, err := s.bannerSvc.Save(r.Context(), item)

	//если получили ошибку то отвечаем с ошибкой
	if err != nil {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusInternalServerError)
		return
	}

	//преобразуем данные в JSON
	data, err := json.Marshal(banner)

	//если получили ошибку то отвечаем с ошибкой
	if err != nil {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusInternalServerError)
		return
	}
	//вызываем функцию для ответа в формате JSON
	respondJSON(w, data)
}

func (s *Server) handleRemoveByID(w http.ResponseWriter, r *http.Request) {

	//извлекаем из параметра запроса ID
	idP := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(idP, 10, 64)
	//если получили ошибку то отвечаем с ошибкой
	if err != nil {

		//печатаем ошибку
		log.Print(err)

		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusBadRequest)
		return
	}

	banner, err := s.bannerSvc.RemoveByID(r.Context(), id)
	//если получили ошибку то отвечаем с ошибкой
	if err != nil {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusInternalServerError)
		return
	}

	//преобразуем данные в JSON
	data, err := json.Marshal(banner)

	//если получили ошибку то отвечаем с ошибкой
	if err != nil {
		//печатаем ошибку
		log.Print(err)
		//вызываем фукцию для ответа с ошибкой
		errorWriter(w, http.StatusInternalServerError)
		return
	}
	//вызываем функцию для ответа в формате JSON
	respondJSON(w, data)
}

/*
+
+
+
+
+
+
+
*/
//это фукция для записывание ошибки в responseWriter или просто для ответа с ошиками
func errorWriter(w http.ResponseWriter, httpSts int) {
	http.Error(w, http.StatusText(httpSts), httpSts)
}

//это функция для ответа в формате JSON
func respondJSON(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(data)
	if err != nil {
		//печатаем ошибку
		log.Print(err)
	}
}
