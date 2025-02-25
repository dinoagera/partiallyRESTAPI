package store_test

import (
	"restapi/internal/app/model"
	"restapi/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepositoriy_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepositoriy_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	email := "example@mail.ru"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
