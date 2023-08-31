package apiserver

import (
	"avitoTestTask/internal/app/model"
	"avitoTestTask/internal/app/tracker"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type server struct {
	router  *mux.Router
	logger  *logrus.Logger
	tracker tracker.Tracker
}

func newServer(tracker tracker.Tracker) *server {
	s := &server{
		router:  mux.NewRouter(),
		logger:  logrus.New(),
		tracker: tracker,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/segment/{slug}", s.handleSegmentCreate()).Methods(http.MethodPost)
	s.router.HandleFunc("/segment/{slug}", s.handleSegmentDelete()).Methods(http.MethodDelete)
	s.router.HandleFunc("/user/{id}", s.handleSegmentsByUserGet()).Methods(http.MethodGet)
	s.router.HandleFunc("/user/{id}", s.handleSegmentsAddToUser()).Methods(http.MethodPost)
	s.router.HandleFunc("/user/{id}", s.handleSegmentsDeleteAtUser()).Methods(http.MethodDelete)
}

func (s *server) handleSegmentsByUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		segments, err := s.tracker.UserSegment().FindByUserId(userId)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, segments)
	}
}

func (s *server) handleSegmentsAddToUser() http.HandlerFunc {
	type request struct {
		Slugs []string `json:"slugs"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		var request request
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&request)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		results := make([]error, 0)
		for _, slug := range request.Slugs {
			userSegment := &model.UserSegment{
				UserId:      userId,
				SegmentSlug: slug,
			}
			err = s.tracker.UserSegment().Create(userSegment)
			results = append(results, err)
		}
		s.respond(w, r, http.StatusCreated, results)
	}
}

func (s *server) handleSegmentsDeleteAtUser() http.HandlerFunc {
	type request struct {
		Slugs []string `json:"slugs"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		var request request
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&request)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		results := make([]error, 0)
		for _, slug := range request.Slugs {
			err = s.tracker.UserSegment().Delete(userId, slug)
			results = append(results, err)
		}
		s.respond(w, r, http.StatusCreated, results)
	}
}

func (s *server) handleSegmentDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slug := mux.Vars(r)["slug"]
		segment := &model.Segment{}

		if err := s.tracker.Segment().Delete(segment, slug); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
		s.respond(w, r, http.StatusOK, segment)
	}
}
func (s *server) handleSegmentCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slug := mux.Vars(r)["slug"]

		segment := &model.Segment{
			Slug: slug,
		}
		if err := s.tracker.Segment().Create(segment); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, segment)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
