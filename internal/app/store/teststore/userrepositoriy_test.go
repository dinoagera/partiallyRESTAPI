package teststore_test

import (
	"restapi/internal/app/model"
	"restapi/internal/app/store"
	"restapi/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepositoriy_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}
func TestUserRepositoriy_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "example@mail.ru"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrorNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
