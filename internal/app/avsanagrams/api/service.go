package api

import (
	"encoding/json"
	"log"
	"net/http"

	"avsanagrams/internal/pkg/anagram"
	"avsanagrams/internal/pkg/storage"
)

type Service struct {
	store storage.Storage
}

func New(st storage.Storage) *Service {
	service := &Service{
		store: st,
	}
	return service
}

func (s *Service) Load(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var words []string

	jd := json.NewDecoder(r.Body)
	if err := jd.Decode(&words); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, w := range words {
		key := anagram.BaseAnagram(w)
		s.store.Add(key, w)
	}

	rw.WriteHeader(http.StatusOK)
}

func (s *Service) Get(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	const queryParamWord = "word"

	word := r.URL.Query().Get(queryParamWord)
	if len(word) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	key := anagram.BaseAnagram(word)
	respBody, err := json.Marshal(s.store.Get(key))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	if _, err := rw.Write(respBody); err != nil {
		log.Fatal("unable to write a response")
	}
}
