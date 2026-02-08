package entity_test

import (
	"testing"

	"github.com/itpaulin/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := entity.NewProduct("Telefone", 999)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.Name, "Telefone")
	assert.Equal(t, p.Price, 999.0)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.CreatedAt)
}

func TestProductWhenNameIsInvalid(t *testing.T) {
	p, err := entity.NewProduct("", 11)
	assert.NotNil(t, p)
	assert.ErrorIs(t, entity.ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := entity.NewProduct("Banana", 0)
	assert.NotNil(t, p)
	assert.ErrorIs(t, entity.ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := entity.NewProduct("Uva", -10)
	assert.NotNil(t, p)
	assert.ErrorIs(t, entity.ErrInvalidPrice, err)
}

func TestProduct_ValidateNewProduct(t *testing.T) {
	p, err := entity.NewProduct("Telefone", 999)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
