package entity_test

import (
	"testing"

	"github.com/itpaulin/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := entity.NewUser("Paulo Ricardo", "pr@gmail.com", "paulo123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Paulo Ricardo", user.Name)
	assert.Equal(t, "pr@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := entity.NewUser("Pr", "pr@g.com", "1234")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("1234"))
	assert.False(t, user.ValidatePassword("123"))
	assert.NotEqual(t, user.Password, "1234")
}
