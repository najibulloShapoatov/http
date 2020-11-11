package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/najibulloShapoatov/http/pkg/banners"
)

//Server ...
type Server struct {
	mux       *http.ServeMux
	bannerSvc *banners.Service
}

//NewServer ....
func NewServer(m *http.ServeMux, bnrSvc *banners.Service) *Server {
	return &Server{mux: m, bannerSvc: bnrSvc}
}

//ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

//Init ...
func (s *Server) Init() {
	s.mux.HandleFunc("/banners.getAll", s.handleGetAllBanners)
	s.mux.HandleFunc("/banners.getById", s.handleGetBannerByID)
	s.mux.HandleFunc("/banners.save", s.handleSaveBanner)
	s.mux.HandleFunc("/banners.removeById", s.handleRemoveByID)
}

func (s *Server) handleGetAllBanners(w http.ResponseWriter, r *http.Request) {

	banners, err := s.bannerSvc.All(r.Context())
	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusInternalServerError)
	}

	data, err := json.Marshal(banners)

	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusInternalServerError)
	}

	respondJSON(w, data)

}

func (s *Server) handleGetBannerByID(w http.ResponseWriter, r *http.Request) {
	idP := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idP, 10, 64)
	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusBadRequest)
	}

	banner, err := s.bannerSvc.ByID(r.Context(), id)

	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusInternalServerError)
	}

	data, err := json.Marshal(banner)

	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusInternalServerError)
	}

	respondJSON(w, data)
}

func (s *Server) handleSaveBanner(w http.ResponseWriter, r *http.Request) {

	idP := r.URL.Query().Get("id")
	title := r.URL.Query().Get("title")
	content := r.URL.Query().Get("content")
	button := r.URL.Query().Get("button")
	link := r.URL.Query().Get("link")

	id, err := strconv.ParseInt(idP, 10, 64)
	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusBadRequest)
	}
	if title == "" && content == "" && button == "" && link == "" {
		log.Print(err)
		errorWriter(w, http.StatusBadRequest)
	}

	item := &banners.Banner{
		ID:      id,
		Title:   title,
		Content: content,
		Button:  button,
		Link:    link,
	}

	banner, err := s.bannerSvc.Save(r.Context(), item)

	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusInternalServerError)
	}

	data, err := json.Marshal(banner)

	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusInternalServerError)
	}
	respondJSON(w, data)
}

func (s *Server) handleRemoveByID(w http.ResponseWriter, r *http.Request) {
	idP := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idP, 10, 64)
	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusBadRequest)
	}

	banner, err := s.bannerSvc.RemoveByID(r.Context(), id)

	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusInternalServerError)
	}

	data, err := json.Marshal(banner)

	if err != nil {
		log.Print(err)
		errorWriter(w, http.StatusInternalServerError)
	}
	respondJSON(w, data)
}

/*
 #
+
+
+
+
+
+
*/

func errorWriter(w http.ResponseWriter, httpSts int) {
	http.Error(w, http.StatusText(httpSts), httpSts)
}

func respondJSON(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(data)
	if err != nil {
		log.Print(err)
	}
}
