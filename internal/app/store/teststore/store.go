package teststore

import (
	"restapi/internal/app/model"
	"restapi/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() store.Store {
	return &Store{}
}
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}
